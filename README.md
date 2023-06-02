# policy_sdk

This repository contains is a software developer kit for [Flexera Policies](https://docs.rightscale.com/policies/). Policies allow you to automate governance across your multi-cloud environment to increase agility and efficiency while managing security and risk in your organization. Developing [custom policies](https://docs.rightscale.com/policies/getting_started/custom_policy.html) allows the policy engine to interface with almost any publicly available API.

## fpt

`fpt` or [Flexera Policy Tool](https://github.com/flexera-public/policy_sdk/tree/master/cmd/fpt/README.md) is a command-line client to aid in the development of custom policies.

## sdk

The [sdk](https://github.com/flexera-public/policy_sdk/tree/master/sdk) directory provides automatically generated code for interfacing with the [Policy API](https://reference.rightscale.com/governance-policies/), written in the Go language. See [client](https://github.com/flexera-public/policy_sdk/tree/master/client) and cmd/fpt directories for examples of code that uses the SDK.

## License

All source code in this repository is subject to the MIT license, see the
[LICENSE](https://github.com/rightscale/fpt/blob/master/LICENSE) file.
