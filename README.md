# txtmsk
txtmsk is encrypts and compress the plain text.

## Install
```
$ go get github.com/ngc224/txtmsk
```

Mac OS X
```
$ brew install ngc224/txtmsk/txtmsk
```

## Usage
set password

- Mac OS X: Keychain
- Linux: Kernel keyring (login session)

```
$ txtmsk -p
```

to mask
```
$ txtmsk "I am a false phimosis"
lKce3vRDwOBa/H7BoEXcXcyw7ZC7LsVkXtmySIZd/sUxABa+caIvUsBB0YlMRJ0rcA
```

to unmask
```
$ txtmsk -u "lKce3vRDwOBa/H7BoEXcXcyw7ZC7LsVkXtmySIZd/sUxABa+caIvUsBB0YlMRJ0rcA"
I am a false phimosis
```

to inline mask

```
$ txtmsk "I am a <txtmsk>false phimosis</txtmsk>"
I am a <txtmsk>oB/3fFLbJlrBsuNgtmHBzZBzOPs7GBdP0vaT4ysTAae48WihwGkzSkdn</txtmsk>
```

to inline unmask
```
$ txtmsk -u "I am a <txtmsk>oB/3fFLbJlrBsuNgtmHBzZBzOPs7GBdP0vaT4ysTAae48WihwGkzSkdn</txtmsk>"
I am a <txtmsk>false phimosis</txtmsk>
```

stdin

```
$ cat credentials.csv | txtmsk
User Name,Access Key Id,Secret Access Key
"nishida",<txtmsk>GPe/IKjqNH1RXTLpk0cco6R8ddWo8l40TT9D69/p8mltCN+nWeYBqkzBqIHeyhqn</txtmsk>,<txtmsk>tYyft6umtmbrCMy55PO3mUg39LNsqFrc6zbWIUZ4wapkYn5PkGZs0HdadtKBEKsgfpa4+GaLSu9y47P/nyHqvbrP1JE</txtmsk>
```

writing inline tag
```
<txtmsk>String of this place are masked.</txtmsk>
```

### Help

```
Usage: txtmsk [options] text
  -h    this help
  -p    set password
  -u    unmask mode
  -v    show version and exit
```