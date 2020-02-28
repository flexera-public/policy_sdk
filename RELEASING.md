# Releasing `policy_sdk`

This is how to release a new version of `policy_sdk`:

1. Verify that all of the tests pass:
   * Linux, macOS, and Windows: [![Travis CI Build Status](https://travis-ci.com/rightscale/policy_sdk.svg?branch=master)](https://travis-ci.com/rightscale/policy_sdk?branch=master)
2. Make sure the [ChangeLog](https://github.com/rightscale/policy_sdk/blob/master/cmd/fpt/ChangeLog.md) and [README](https://github.com/rightscale/policy_sdk/blob/master/README.md) are up to date.
3. Create a tag of the form `vX.Y.Z` where `X`, `Y`, and `Z` are the major, minor, and patch versions respectively:
   ```bash
   git checkout master
   git pull
   git tag --annotate --message='Release version vX.Y.Z' vX.Y.Z
   git push --tags
   ```
4. Create a [GitHub release](https://github.com/rightscale/policy_sdk/releases) from the tag with the ChangeLog contents
  as the description. Also include links to the binaries for Linux, macOS, and Windows in the description:
   * Linux: `https://binaries.rightscale.com/rsbin/policy_sdk/vX.Y.Z/right_st-linux-amd64.tgz`
   * macOS: `https://binaries.rightscale.com/rsbin/policy_sdk/vX.Y.Z/right_st-darwin-amd64.tgz`
   * Windows: `https://binaries.rightscale.com/rsbin/policy_sdk/vX.Y.Z/right_st-windows-amd64.zip`

## Testing the release

TBD: should include testing binaries for Linux, macOS, and Windows
