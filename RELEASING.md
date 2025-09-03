# Releasing `policy_sdk`

This is how to release a new version of `policy_sdk`:

1. Verify that all of the tests pass:
   * Linux, macOS, and Windows: [![Travis CI Build Status](https://travis-ci.com/rightscale/policy_sdk.svg?branch=master)](https://travis-ci.com/rightscale/policy_sdk?branch=master)
2. Make sure the [ChangeLog](https://github.com/flexera-public/policy_sdk/blob/master/cmd/fpt/ChangeLog.md) and [README](https://github.com/flexera-public/policy_sdk/blob/master/README.md) are up to date.
3. Create a tag of the form `vX.Y.Z` where `X`, `Y`, and `Z` are the major, minor, and patch versions respectively:
   ```bash
   git checkout master
   git pull
   git tag --annotate --message='Release version vX.Y.Z' vX.Y.Z
   git push --tags
   ```
4. Once the [Travis CI tag build](https://travis-ci.com/github/rightscale/policy_sdk/builds) succeeds, run the release script to create a GitHub Release:
   ```bash
   ./release.sh
   ```
   This script uses the [`hub`](https://hub.github.com/) command line tool for GitHub, you will need to install it using the [Installation](https://github.com/github/hub#installation) instructions for your platform and authenticate it with GitHub before running the script (you can use the `hub release` command to test if you are authenticated correctly).

## Testing the release

TBD: should include testing binaries for Linux, macOS, and Windows
