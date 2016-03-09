txtmsk
=======
txtmsk encrypts the plain text.

Install
-----
    $ brew install ngc224/txtmsk/txtmsk

Usage
-----
Set Password

- Mac OS X: Keychain
- Linux: Kernel keyring (login session)


    $ txtmsk -p

Encrypt

    $ txtmsk aaaaaaaaaaaaaaaaaaaa
    GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc

Decrypt

    $ txtmsk -d GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
    aaaaaaaaaaaaaaaaaaaa

Encrypt stdin

    $ cat ~/.aws/credentials | txtmsk > credentials.txtmsk


Decrypt stdin

    $ cat credentials.txtmsk | txtmsk -d > ~/.aws/credentials

Help
-----
    Usage:
      txtmsk [-d] [-p] [-v] text
    
    Application Options:
      -d, --decrypt   Decrypt mode
      -p, --password  Set the password in Keychain
      -v, --version   Show program's version number
    
    Help Options:
      -h, --help      Show this help message