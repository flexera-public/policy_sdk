# Releasing `policy_sdk`

This is how to release a new version of `policy_sdk`:

1. Verify that all of the tests pass:
   * Linux, macOS, and Windows: [![Test Policy SDK](https://github.com/flexera-public/policy_sdk/actions/workflows/test.yml/badge.svg?event=push)](https://github.com/flexera-public/policy_sdk/actions/workflows/test.yml) 
2. Create a tag of the form `vX.Y.Z` where `X`, `Y`, and `Z` are the major, minor, and patch versions respectively:
   ```bash
   git checkout master
   git pull
   git tag --annotate --message='Release version vX.Y.Z' vX.Y.Z
   git push --tags
   ```
3. Once the [GitHub Actions Workflow](https://github.com/flexera-public/policy_sdk/actions/workflows/build.yml) succeeds, review the draft release created under [Releases](https://github.com/flexera-public/policy_sdk/releases)

## Testing the release

TBD: should include testing binaries for Linux, macOS, and Windows
