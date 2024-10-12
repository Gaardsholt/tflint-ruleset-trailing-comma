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
  source  = "github.com/Gaardsholt/tflint-ruleset-trailing-comma"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----

  mQGNBGcKjvABDADOtyLeAVdP/bjqdvukjvOqdQ0q/l1vyWTtilb9haLUFcBAos1Q
  knjyq4Q0XWcs8HdB9lnd7mvd37Tut5D2t4RMlKWbANGgU286WaLdf9P0a62yN1ID
  TdobfcWoJrQgrx5Wx24r4WPOCPjVoW4bYX4zO588WDTXu+OLtJ1d6vNgdtEx6ck+
  oL6eg5nBqya8s3xHbQx0aXWwUHFDTAlHG5UfBoOM2t9ROhdDIF6aCby68piuXBOV
  L4vXmbbfR2vdMZvq7/zFCtER1kpM76To4mElsi9QzmFiRGcPk4DyGUOyrfux9cI/
  IpTbZJMNXjS/lq7l6OU5Mnpijk7vUyewM9o0RKO1KCN0JHzjTX/2AZHhuL6l/6+m
  nj/m3LADZFKck+rd7cQSZPCEHapen+wy6MEbsFChbnSZSDJrDupOfpI8xAE1ndxI
  jhSg9a+qOTKaRvJXbVNwFq2SEHkKfbuH5DwSshg4YG3A+SDQ7mI5fZAyTsMVBqBO
  o5GYqlb0b8cMjK8AEQEAAbQtTGFzc2UgR2FhcmRzaG9sdCA8bGFzc2UuZ2FhcmRz
  aG9sdEBnbWFpbC5jb20+iQHRBBMBCAA7FiEEyLpGOfaGJKmkCXNOsgSkCsdnwOwF
  AmcKjvACGwMFCwkIBwICIgIGFQoJCAsCBBYCAwECHgcCF4AACgkQsgSkCsdnwOze
  Cgv9H2WoSDMhamZlCDx87+5rrQ43EraTMgh612Gtj8Grbpd5T1alHfN99N22yqbk
  oiGH37kBXADbxPbvPruH093HVaoR3u05tfEDHHU5v4Vjw1TpBlEkQD/6OQHkFsCn
  Dn02CXk1r6Jcc6AqKDrNwM4nRxAZutdcyRplrA6WFNXp4RxeA2wyW4dl9gEkplJj
  /R8NGkLlUNUII8dtLEeamKzVj+zUaglIWBvc9OVlf//dqtI7sU5fLxHlrmDRFVyF
  M+sZNfccvbwbAZ4BQHxFxAayfMd05PREeBjZfMkBjooyN3HgSJmhMCTucK6JP7PA
  cERIC30oFmpfAEarV/nBqkgF4hrJTxUbKbIGXbMaR7/aMrVTaT+e7KNe84p95bOg
  yh+ROR4qejGendH5EjR0t7JuYVMS2v5UvLAc8ENhNUdt+6bRC34rSnVDcKqcWFab
  cbb/oevglQTj9LdwSBolzoQaNhrttgR3aqMUQcCPXG75VWCoAzQwSqbMHOrs0DPA
  EMhguQGNBGcKjvABDACeDrPJz2w1Cnl/FddB/JvGCCl8Pl+wW/+w105uooPfhZVg
  6ypGcvzLIG+VJP9FTEq7/yUcQRnMlB6BQXKCE/3MxICY8Srt7Q3rYXZEKT39ox2E
  zSYb1oXrAtSHvyF8eP7mRwESvCkvGxQHD0IJUUilNrVXcszccE9gwSMv29lrDO8M
  6iAKqRY9Oqrwn7rJws+RbbIRv1dPDpGq4EVK3vSQUB/ORetKyky4YBi8s3z0LRr9
  wyTTh89dmwlRM3Pfmnx2jn9M5UBNV3waSZCZoR1cFwKGHjDbizx3uPQCaj6YRn80
  qDTqYXvPcS+rYR/kd4OkBmVjc3k2szWv+E2Shch3H4q2177uMNLAFL7jLOHRjGba
  8kdKCwuK5kgxpIP6lEkPG7MCsLDp3DQ2srxmLR68wyrXIiycnlEXKXkTJTR9gTFK
  foc43Mk8EfsE0V7lfV35+l+M9fyKQDLXPa59cyYVSbd8vX38PTsGk+CoNSslbaFC
  l4mrJKD1Lj29xSTQyWkAEQEAAYkBtgQYAQgAIBYhBMi6Rjn2hiSppAlzTrIEpArH
  Z8DsBQJnCo7wAhsMAAoJELIEpArHZ8Ds2HoL/iU0FeYPjB3N8JngXLyklc5ZuDC7
  fGLTEqHO31bGhB7Bs+7xojgwrK+zvswGJU+ByoLUHP++SVal1lZ8OwbygM46NGjj
  F5OKVJ2MYt8GGwPeSrxhrAokxH7sAleENCy+IuwoD8e6pFr53o9KXCNGIOp65SM6
  /RuBo+Tt/P8lz4TibSMCsP+zo1q0yhHo2kOvDVqnFGQGbiRI6PCdotMT0c9bkHxa
  JGG2fNbiriIf6cCV25G7Ajf8OCYhiQGhDyZYpeozfHZJORFxd1z+VsdW4RX2zVIw
  B4OT9Sg2ujkrdgnHgCAa36PS9jDel+z04DqL5gVYPjXpR2vdLQORZgJwtP8XY0g/
  NRZWg5k99kHCLBDg7kLjDJpx7UHE7+Q98Fxi/5HDYrBQ+dHcDiHpIZuiIaUAG45v
  csie3rKvoDiXBOccjXrxJNbhqN6RTuqBts+atNgz7Mn4j6YGCIoF03gfsxGc51ZB
  haUXwAPe9pvS8sCURInYGfMFlThmvN5ZuYkJ3g==
  =U3wP
  -----END PGP PUBLIC KEY BLOCK-----
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
