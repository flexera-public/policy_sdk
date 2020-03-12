#!/usr/bin/env bash

set -e

# use GNU coreutils sort because it supports version sort
case `uname` in
(Darwin)
  sed='gsed'
  sort='gsort'
  ;;
(*)
  sed='sed'
  sort='sort'
  ;;
esac

# find the latest version tagged in Git
while read version; do
  if [[ $version =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    break
  fi
done < <(git tag -l 'v*' | $sort --reverse --version-sort)

# create a GitHub release from the latest version tagged in Git with ChangeLog entries and pre-compiled binary links
hub release create --file - $version <<EOF
$version Release

$($sed -nre "/^${version//./\\.} \/ .+\$/,/^$/{/^${version//./\\.} \/ .+\$/d;/^-+$/d;p}" cmd/fpt/ChangeLog.md)

Pre-compiled binaries:
* Linux: [$version/fpt-linux-amd64.tgz](https://binaries.rightscale.com/rsbin/fpt/$version/fpt-linux-amd64.tgz)
* macOS: [$version/fpt-darwin-amd64.tgz](https://binaries.rightscale.com/rsbin/fpt/$version/fpt-darwin-amd64.tgz)
* Windows: [$version/fpt-windows-amd64.zip](https://binaries.rightscale.com/rsbin/fpt/$version/fpt-windows-amd64.zip)
EOF
