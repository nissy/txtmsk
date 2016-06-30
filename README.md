# txtmsk
txtmsk is encrypts and compress the plain text.

encryption is aes256

compress is zlib

## Install
```
$ go get github.com/ngc224/txtmsk
```

Mac OS X
```
$ brew install ngc224/txtmsk/txtmsk
```

## Usage
#### Set password

- Mac OS X: Keychain
- Linux: Kernel keyring (login session)

```
$ txtmsk -p
```

#### Mask & Unmask

Text to mask
```
$ txtmsk 'I am a false phimosis'
lKce3vRDwOBa/H7BoEXcXcyw7ZC7LsVkXtmySIZd/sUxABa+caIvUsBB0YlMRJ0rcA
```

Text to unmask
```
$ txtmsk -u 'lKce3vRDwOBa/H7BoEXcXcyw7ZC7LsVkXtmySIZd/sUxABa+caIvUsBB0YlMRJ0rcA'
I am a false phimosis
```

Text to mask (stdin)
```
$ echo 'I am a false phimosis' | txtmsk
lKce3vRDwOBa/H7BoEXcXcyw7ZC7LsVkXtmySIZd/sUxABa+caIvUsBB0YlMRJ0rcA
```

Text to unmask (stdin)
```
$ echo 'lKce3vRDwOBa/H7BoEXcXcyw7ZC7LsVkXtmySIZd/sUxABa+caIvUsBB0YlMRJ0rcA' | txtmsk -u
I am a false phimosis
```

#### Inline Mask & Inline Unmask

Writing inline tag
```
<msk>text</msk>
```

Text to inline mask
```
$ cat secret.txt
I am a <msk>false phimosis</msk>
I am a <msk>false phimosis</msk>
I am a <msk>false phimosis</msk>
I am a <msk>false phimosis</msk>
I am a <msk>false phimosis</msk>

$ cat secret.txt | txtmsk
I am a <msk>n6kL28dZQURtJ3as/Hpsryp+OwBR2rAN3Dbgb3iT84Mz7/f3gIu7qhqF</msk>
I am a <msk>ruM8Rs1otjSrp5UhIOd5Z7Et6eC3zdBDlX3UaLvrPBAS0Hm6mOnZ1zjr</msk>
I am a <msk>9NbYFdyGKycism9hx5Pq1hwGLNxz9+89Y02IL5ux9Nwt0QaUQGZKMeVS</msk>
I am a <msk>ik02zV8PA2QxXV379KV0KRCVastEoJNVqkqEHyTKrb45Y05Rd142cQJn</msk>
I am a <msk>OK44PxypCUO7KvYH+U8iyaYRzvaoqcTh8yHMkvcenUNP+6seRvVgWLP8</msk>
```

Text to inline unmask ("-t" option is trim inline tags)
```
$ cat secret.txtmsk.txt
I am a <msk>n6kL28dZQURtJ3as/Hpsryp+OwBR2rAN3Dbgb3iT84Mz7/f3gIu7qhqF</msk>
I am a <msk>ruM8Rs1otjSrp5UhIOd5Z7Et6eC3zdBDlX3UaLvrPBAS0Hm6mOnZ1zjr</msk>
I am a <msk>9NbYFdyGKycism9hx5Pq1hwGLNxz9+89Y02IL5ux9Nwt0QaUQGZKMeVS</msk>
I am a <msk>ik02zV8PA2QxXV379KV0KRCVastEoJNVqkqEHyTKrb45Y05Rd142cQJn</msk>
I am a <msk>OK44PxypCUO7KvYH+U8iyaYRzvaoqcTh8yHMkvcenUNP+6seRvVgWLP8</msk>

$ cat secret.txtmsk.txt | txtmsk -u
I am a <msk>false phimosis</msk>
I am a <msk>false phimosis</msk>
I am a <msk>false phimosis</msk>
I am a <msk>false phimosis</msk>
I am a <msk>false phimosis</msk>

$ cat secret.txtmsk.txt | txtmsk -u -t
I am a false phimosis
I am a false phimosis
I am a false phimosis
I am a false phimosis
I am a false phimosis
```

### Help

```
Usage: txtmsk [options] text
  -h    this help
  -p    set password
  -t    trim inline tags (unmask mode only)
  -u    unmask mode
  -v    show version and exit
```