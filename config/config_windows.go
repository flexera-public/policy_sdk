package config

import (
	"path/filepath"

	"github.com/douglaswth/rsrdp/win32"
)

func DefaultFile(name string) string {
	roamingPath, err := win32.SHGetKnownFolderPath(&win32.FOLDERID_RoamingAppData, 0, 0)
	if err != nil {
		panic(err)
	}
	return filepath.Join(roamingPath, "RightScale", name)
}
