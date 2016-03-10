# txtmsk
txtmsk encrypts the plain text.

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

### Encrypt
```
$ txtmsk aaaaaaaaaaaaaaaaaaaa
GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
```

or

```
$ echo aaaaaaaaaaaaaaaaaaaa | txtmsk
GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
```

### Decrypt
```
$ txtmsk -d GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
aaaaaaaaaaaaaaaaaaaa
```

or

```
$ echo GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc | txtmsk -d
aaaaaaaaaaaaaaaaaaaa
```

### Help

```
Usage:
  txtmsk [-d] [-p] [-v] text

Application Options:
  -d, --decrypt   Decrypt mode
  -p, --password  Set the password
  -v, --version   Show program's version number

Help Options:
  -h, --help      Show this help message
```