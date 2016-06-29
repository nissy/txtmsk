# txtmsk
txtmsk encrypts and compress the plain text.

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
$ txtmsk aaaaaaaaaaaaaaaaaaaa
GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
```

or

```
$ echo aaaaaaaaaaaaaaaaaaaa | txtmsk
GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
```

### UnMask
```
$ txtmsk -u GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
aaaaaaaaaaaaaaaaaaaa
```

or

```
$ echo GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc | txtmsk -u
aaaaaaaaaaaaaaaaaaaa
```

### Help

```
Usage: txtmsk [options] text
  -h    this help
  -p    set password
  -u    unmask mode
  -v    show version and exit
```