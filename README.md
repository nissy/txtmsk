# txtmsk
txtmsk is encrypts and compress the plain text.

## Install
```
$ brew install ngc224/txtmsk/txtmsk
```

## Usage
### Set Password

- Mac OS X: Keychain
- Linux: Kernel keyring (login session)

```
$ txtmsk -p
```

### Mask
```
$ txtmsk aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
mtf1fFlPw9NU7lH56c01hgtD+PAp94SYY7rbSq5HYQ==
```

or

```
$ echo aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa | txtmsk
mtf1fFlPw9NU7lH56c01hgtD+PAp94SYY7rbSq5HYQ==
```

### Unmask
```
$ txtmsk -u mtf1fFlPw9NU7lH56c01hgtD+PAp94SYY7rbSq5HYQ==
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
```

or

```
$ echo mtf1fFlPw9NU7lH56c01hgtD+PAp94SYY7rbSq5HYQ== | txtmsk -u
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
```

### Help

```
Usage: txtmsk [options] text
  -h    this help
  -p    set password
  -u    unmask mode
  -v    show version and exit
```