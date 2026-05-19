# zal-go

z̢͎͍̱̬͙̭͙̻̜͕̝̮̄̾̉͊̉̋̅̋̒͆̑̄̂̃̓͟͜ą̶̸̛̛̤̹̩̟̤̜̮͙̟͍̰̐̽͋̆́͆̕͟͠͠l̵̡̛̪̮̘̱̻̪̗̱̤̰͙͕̩̆̌̽̓̆̓̊͆̍̓́͌̉͞-̵̸̨̨̧̛͖̟̝͚͋̃̎̀̂̑̒͢͞͡g̷̢̢̡̭͎̭̘͚̱̘̩̪͖̝͍͖̑̊̏̎͞ő̧̻̟̭̫̹̦͕͍̰̾̄͊̒̎̈́̍̐̉̆́̕
## Description

This is a simple CLI utility to make text appear corrupted.

## Usage

```console
$ echo "Hello, world!" | zalgo
H̞͖̚̚e̼l̶̬̩l͙̱õ͓̬,̜̭̬̭ ̑w̢͖͌̾o̲̻r͚ļ͍d̦͎̂!̀̾͠
```

The severity of the corruption can be adjusted with the `--min` and `--max` options.

## Installation

### From GitHub Releases

Use the appropriate script, or just download it from the releases page and put it somewhere on `PATH`.

#### Unix

*You may need to call `sudo sh` instead.*

```shell
curl -fsSL https://github.com/spenserblack/zal-go/raw/refs/heads/main/_scripts/install-gh.sh | sh
```

### `go install`

*This installation method will install the executable as `zal-go`.*

```shell
go install github.com/spenserblack/zal-go@latest
```
