# zal-go

z̢͎͍̱̬͙̭͙̻̜͕̝̮̄̾̉͊̉̋̅̋̒͆̑̄̂̃̓͟͜ą̶̸̛̛̤̹̩̟̤̜̮͙̟͍̰̐̽͋̆́͆̕͟͠͠l̵̡̛̪̮̘̱̻̪̗̱̤̰͙͕̩̆̌̽̓̆̓̊͆̍̓́͌̉͞-̵̸̨̨̧̛͖̟̝͚͋̃̎̀̂̑̒͢͞͡g̷̢̢̡̭͎̭̘͚̱̘̩̪͖̝͍͖̑̊̏̎͞ő̧̻̟̭̫̹̦͕͍̰̾̄͊̒̎̈́̍̐̉̆́̕
## Description

This is a simple CLI utility to make text appear corrupted.

## Usage

### Web

Go to [spenserblack.github.io/zal-go](https://spenserblack.github.io/zal-go/).

### CLI

```console
$ echo "Hello, world!" | zalgo
H̞͖̚̚e̼l̶̬̩l͙̱õ͓̬,̜̭̬̭ ̑w̢͖͌̾o̲̻r͚ļ͍d̦͎̂!̀̾͠

$ zalgo progressive "This is a very very very long string that gets progressively more corrupt."
T͖his̎͜ ̣̇i͓̳s̮̟̎ ̺̰́ȁ̺͆ ̩̖v̈́̓̈e͂̍r̠͓͜y͓̻̎ ̸̜̫͎͂v͍͓̓̆̏ě̬̀r͙͟ỳ̛̥̮̅͢ ̧̫͌̂̋̈v̭̻̻̓ę̗̟̬̯̀r̞͖̯̻̔͢͝y̥̭͋̆ ̇̀̇͢l̛̫͚̼͍̽̈́͠o̸͕̓̍̎͆͌̕n̸̮̋͡g͙̗̍̌̓̀ ̸̲̥͂s̯͙̬̽t̨̧̞̺̎͌̉͢r̡̮̱̼̐̑ḯ͓̬͔̏ň̶̛̞̜͡͡g̯͖̲̥̗̪̖̀̃̀ ̷̼̎͋̂̀̈́̌t̸̢̳̞̬͆h̖̱͜͟͝a̺̬̖͚͝t̙͎̤̱͔͌̃̒̓͞ ̛͕̻̟̇̃̀̎̂̈́g̹̠̙̖̝̺͎͚͋̏̉e̷̫̜̯̼̽̔͝ẗ̫͖͙̏͊s̜͋̄̽̕͟ ̖͖̗̎̋̊͜͞p̢̢̱̩̬̊̑r̨̟̮͔̬̐̓͊̆ǭ̢̥̬̦̭͊͊̏͆̃̋̐g̸̭̗̭̾͠͞r̛̩̘̭̱̅̓̑̿͆͆̚e̸͔̯̰̎̒́͡s̷̢̗̯̹̠̬̟̽͋̉͝s̳̭̭̄̍̍̆̆̅̅̚ī̛̮̻̳͋͌͢͠v̶̡̜̭̻̠̱̪̠̟͞ê̸̡̙̳͖̱̠̆̕͜l͕͔̤͙̭̫͕̺̀͆́̔ý̵̢̛̠̭̫̝̜̫͙̙̍̌͌̍ ̷̧̥̘̮͔̤͚͎̊̄͂̑̿͜͜͠m͎͍͕̪͙̓̏̐̅̀̕͞ò͔̦͔̙̗̀̀͜ṟ̱͚̐̑̊͌͜͡ę̷̥̜̗̱̜̈́̋̅͌́̕ ̧͕͚͓̹̼̳͓̯̀̔̎̉͡͝͠ć̼̬͔̳̄̾̎͟õ̲͎̦̱̭̩̔̉͡ŗ̰̣̹̐̋̍̍͆̚ṙ̢̺̦̳̑̒̋̋͢ṳ̬̭̙̥̻̠̃̓̓̿̑̚͢͟͟͝p̡̫̜̑̅̑̕͡͡͝t̢̛̹̤̼̟̺͖̣͙͕̊̓.̷̨̢̜̥̠̩͕͚̙̙͖̏̿̅̓̌͜͝͝
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
