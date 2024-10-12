# TFLint Ruleset Template
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-template/workflows/build/badge.svg?branch=main)](https://github.com/terraform-linters/tflint-ruleset-template/actions)

This is a template repository for building a custom ruleset. You can create a plugin repository from "Use this template". See also [Writing Plugins](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/plugins.md).

## Requirements

- TFLint v0.42+
- Go v1.22

## Installation

TODO: This template repository does not contain release binaries, so this installation will not work. Please rewrite for your repository. See the "Building the plugin" section to get this template ruleset working.

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "template" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/terraform-linters/tflint-ruleset-template"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----
  mQINBGCqS2YBEADJ7gHktSV5NgUe08hD/uWWPwY07d5WZ1+F9I9SoiK/mtcNGz4P
  JLrYAIUTMBvrxk3I+kuwhp7MCk7CD/tRVkPRIklONgtKsp8jCke7FB3PuFlP/ptL
  SlbaXx53FCZSOzCJo9puZajVWydoGfnZi5apddd11Zw1FuJma3YElHZ1A1D2YvrF
  ...
  KEY
}
```

## Rules

| Name                           | Description                                            | Severity | Enabled | Link |
| ------------------------------ | ------------------------------------------------------ | -------- | ------- | ---- |
| terraform_lists_trailing_comma | Will check if last item in a list has a trailing comma | ERROR    | ✔       |      |

## Building the plugin

Clone the repository locally and run the following command:

```shell
make
```

You can easily install the built plugin with the following:

```shell
make install
```

You can run the built plugin like the following:

```shell
cat << EOS > .tflint.hcl
plugin "trailing-comma" {
  enabled = true
}
EOS
tflint
```
