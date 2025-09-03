package main_test

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	. "github.com/flexera-public/policy_sdk/cmd/fpt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Update", func() {
	Describe("Get current version", func() {
		It("Gets a version from a tagged version", func() {
			version := UpdateGetCurrentVersion("fpt v98.76.54 - JUNK JUNK JUNK")
			Expect(version).To(Equal(&Version{98, 76, 54}))
		})

		It("Gets no version for a dev version", func() {
			version := UpdateGetCurrentVersion("fpt master - JUNK JUNK JUNK")
			Expect(version).To(BeNil())
		})
	})

	Context("With a update versions URL", func() {
		var (
			buffer                *gbytes.Buffer
			server                *httptest.Server
			newExeContent         string
			oldUpdateGithubAPIUrl string
		)

		BeforeEach(func() {
			buffer = gbytes.NewBuffer()
			exeItem := "fpt/fpt"
			if runtime.GOOS == "windows" {
				exeItem += ".exe"
			}
			ext := "tgz"
			if runtime.GOOS == "windows" {
				ext = "zip"
			}
			assetPattern := regexp.MustCompile(`^/fpt-` + runtime.GOOS + `-` + runtime.GOARCH + `\.` + ext + `$`)

			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/":
					// Mock GitHub releases API response
					w.Header().Set("Content-Type", "application/json")
					assetName := "fpt-" + runtime.GOOS + "-" + runtime.GOARCH + "." + ext
					downloadURL := server.URL + "/" + assetName
					response := `[
						{
							"tag_name": "v3.4.5",
							"draft": false,
							"prerelease": false,
							"assets": [
								{
									"name": "` + assetName + `",
									"browser_download_url": "` + downloadURL + `"
								}
							]
						},
						{
							"tag_name": "v2.3.4",
							"draft": false,
							"prerelease": false,
							"assets": [
								{
									"name": "` + assetName + `",
									"browser_download_url": "` + downloadURL + `"
								}
							]
						},
						{
							"tag_name": "v1.2.3",
							"draft": false,
							"prerelease": false,
							"assets": [
								{
									"name": "` + assetName + `",
									"browser_download_url": "` + downloadURL + `"
								}
							]
						}
					]`
					w.Write([]byte(response))
				default:
					if assetPattern.MatchString(r.URL.Path) {
						// Serve the mock binary archive
						if ext == "tgz" {
							gzipWriter := gzip.NewWriter(w)
							tarWriter := tar.NewWriter(gzipWriter)
							if err := tarWriter.WriteHeader(&tar.Header{Name: exeItem, Size: int64(len(newExeContent))}); err != nil {
								panic(err)
							}
							if _, err := io.WriteString(tarWriter, newExeContent); err != nil {
								panic(err)
							}
							if err := tarWriter.Close(); err != nil {
								panic(err)
							}
							if err := gzipWriter.Close(); err != nil {
								panic(err)
							}
						} else {
							zipWriter := zip.NewWriter(w)
							exeWriter, err := zipWriter.Create(exeItem)
							if err != nil {
								panic(err)
							}
							if _, err := io.WriteString(exeWriter, newExeContent); err != nil {
								panic(err)
							}
							if err := zipWriter.Close(); err != nil {
								panic(err)
							}
						}
					} else {
						w.WriteHeader(http.StatusNotFound)
					}
				}
			}))
			newExeContent = "#!/bin/bash\necho 'This is the new version!'\n"
			oldUpdateGithubAPIUrl = UpdateGithubAPIUrl
			UpdateGithubAPIUrl = server.URL
		})

		AfterEach(func() {
			UpdateGithubAPIUrl = oldUpdateGithubAPIUrl
			server.Close()
		})

		Describe("Get latest versions", func() {
			It("Gets the latest versions", func() {
				latest, err := UpdateGetLatestVersions()
				Expect(err).NotTo(HaveOccurred())
				Expect(latest).To(Equal(&LatestVersions{
					Versions: map[int]*Version{
						1: &Version{1, 2, 3},
						2: &Version{2, 3, 4},
						3: &Version{3, 4, 5},
					},
				}))
			})
		})

		Describe("Get download URL", func() {
			var ext string
			if runtime.GOOS == "windows" {
				ext = "zip"
			} else {
				ext = "tgz"
			}

			It("Gets the download URL for a major version", func() {
				url, version, err := UpdateGetDownloadUrl(1)
				Expect(err).NotTo(HaveOccurred())
				Expect(url).To(Equal(server.URL + "/fpt-" + runtime.GOOS + "-" + runtime.GOARCH + "." + ext))
				Expect(version).To(Equal(&Version{1, 2, 3}))
			})

			It("Returns an error for a nonexistent major version", func() {
				url, version, err := UpdateGetDownloadUrl(0)
				Expect(err).To(MatchError("major version not available: 0"))
				Expect(url).To(BeEmpty())
				Expect(version).To(BeNil())
			})
		})

		Describe("Update check", func() {
			It("Outputs nothing for a dev version", func() {
				UpdateCheck("fpt dev - JUNK JUNK JUNK", buffer)
				Expect(buffer.Contents()).To(BeEmpty())
			})

			It("Outputs nothing if there is no update", func() {
				UpdateCheck("fpt v3.4.5 - JUNK JUNK JUNK", buffer)
				Expect(buffer.Contents()).To(BeEmpty())
			})

			It("Outputs that there is a new version", func() {
				UpdateCheck("fpt v3.0.0 - JUNK JUNK JUNK", buffer)
				Expect(buffer.Contents()).To(BeEquivalentTo(`There is a new v3 version of fpt (v3.4.5), to upgrade run:
    fpt update apply

See https://github.com/flexera-public/policy_sdk/releases for more information.
`))
			})

			It("Outputs that there is a new major version", func() {
				UpdateCheck("fpt v2.3.4 - JUNK JUNK JUNK", buffer)
				Expect(buffer.Contents()).To(BeEquivalentTo(`There is a new major version of fpt (v3.4.5), to upgrade run:
    fpt update apply -m 3

See https://github.com/flexera-public/policy_sdk/releases for more information.
`))
			})

			It("Outptus that there is a new version and new major version", func() {
				UpdateCheck("fpt v2.0.0 - JUNK JUNK JUNK", buffer)
				Expect(buffer.Contents()).To(BeEquivalentTo(`There is a new v2 version of fpt (v2.3.4), to upgrade run:
    fpt update apply
There is a new major version of fpt (v3.4.5), to upgrade run:
    fpt update apply -m 3

See https://github.com/flexera-public/policy_sdk/releases for more information.
`))
			})
		})

		Describe("Update list", func() {
			It("Outputs the available versions for a dev version", func() {
				Expect(UpdateList("fpt dev - JUNK JUNK JUNK", buffer)).To(Succeed())
				Expect(buffer.Contents()).To(BeEquivalentTo(`The latest v1 version of fpt is v1.2.3.
The latest v2 version of fpt is v2.3.4.
The latest v3 version of fpt is v3.4.5.

See https://github.com/flexera-public/policy_sdk/releases for more information.
`))
			})

			It("Outputs the available versions for an up to date version", func() {
				Expect(UpdateList("fpt v3.4.5 - JUNK JUNK JUNK", buffer)).To(Succeed())
				Expect(buffer.Contents()).To(BeEquivalentTo(`The latest v1 version of fpt is v1.2.3.
The latest v2 version of fpt is v2.3.4.
The latest v3 version of fpt is v3.4.5; this is the version you are using!

See https://github.com/flexera-public/policy_sdk/releases for more information.
`))
			})

			It("Outputs the available versions when there is a new version", func() {
				Expect(UpdateList("fpt v3.0.0 - JUNK JUNK JUNK", buffer)).To(Succeed())
				Expect(buffer.Contents()).To(BeEquivalentTo(`The latest v1 version of fpt is v1.2.3.
The latest v2 version of fpt is v2.3.4.
The latest v3 version of fpt is v3.4.5; you are using v3.0.0, to upgrade run:
    fpt update apply

See https://github.com/flexera-public/policy_sdk/releases for more information.
`))
			})

			It("Outputs the available versions when there is a new major version", func() {
				Expect(UpdateList("fpt v2.3.4 - JUNK JUNK JUNK", buffer)).To(Succeed())
				Expect(buffer.Contents()).To(BeEquivalentTo(`The latest v1 version of fpt is v1.2.3.
The latest v2 version of fpt is v2.3.4; this is the version you are using!
The latest v3 version of fpt is v3.4.5; you are using v2.3.4, to upgrade run:
    fpt update apply -m 3

See https://github.com/flexera-public/policy_sdk/releases for more information.
`))
			})

			It("Outputs the available versions when there is a new version and new major version", func() {
				Expect(UpdateList("fpt v2.0.0 - JUNK JUNK JUNK", buffer)).To(Succeed())
				Expect(buffer.Contents()).To(BeEquivalentTo(`The latest v1 version of fpt is v1.2.3.
The latest v2 version of fpt is v2.3.4; you are using v2.0.0, to upgrade run:
    fpt update apply
The latest v3 version of fpt is v3.4.5; you are using v2.0.0, to upgrade run:
    fpt update apply -m 3

See https://github.com/flexera-public/policy_sdk/releases for more information.
`))
			})
		})

		Describe("Update apply", func() {
			var (
				tempDir string
				exePath string
			)

			BeforeEach(func() {
				var err error
				tempDir, err = os.MkdirTemp("", "update")
				if err != nil {
					panic(err)
				}
				exePath = filepath.Join(tempDir, "fpt")
				err = os.WriteFile(exePath, []byte("#!/bin/bash\necho 'This is the old version!'\n"), 0755)
				if err != nil {
					panic(err)
				}
			})

			AfterEach(func() {
				os.RemoveAll(tempDir)
			})

			It("Updates to the latest version for the current major version", func() {
				Expect(UpdateApply("fpt v2.0.0 - JUNK JUNK JUNK", buffer, 0, exePath)).To(Succeed())
				ext := "tgz"
				if runtime.GOOS == "windows" {
					ext = "zip"
				}
				Expect(buffer.Contents()).To(MatchRegexp(`^Downloading fpt v2\.3\.4 from %s/fpt-%s-%s\.%s\.\.\.
Successfully updated fpt to v2\.3\.4!
$`, regexp.QuoteMeta(server.URL), runtime.GOOS, runtime.GOARCH, ext))

				exeContent, err := os.ReadFile(exePath)
				Expect(err).NotTo(HaveOccurred())
				Expect(exeContent).To(BeEquivalentTo(newExeContent))
			})

			It("Updates to the latest version for a specific major version", func() {
				Expect(UpdateApply("fpt v2.0.0 - JUNK JUNK JUNK", buffer, 3, exePath)).To(Succeed())
				ext := "tgz"
				if runtime.GOOS == "windows" {
					ext = "zip"
				}
				Expect(buffer.Contents()).To(MatchRegexp(`^Downloading fpt v3\.4\.5 from %s/fpt-%s-%s\.%s\.\.\.
Successfully updated fpt to v3\.4\.5!
$`, regexp.QuoteMeta(server.URL), runtime.GOOS, runtime.GOARCH, ext))

				exeContent, err := os.ReadFile(exePath)
				Expect(err).NotTo(HaveOccurred())
				Expect(exeContent).To(BeEquivalentTo(newExeContent))
			})
		})
	})
})
