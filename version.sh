#!/usr/bin/env bash

set -e

# hardcoding this part for backwards compatibility, version releases going forward will use github's release tagging feature
cat <<EOF
# Latest fpt versions by major version (this file is used by fpt's update check mechanism)
---
versions:
  1: $GITHUB_REF_NAME
EOF
