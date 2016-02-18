txtmsk
=======
txtmsk encrypts the plain text.

Mac OS X only !!!


Install
-----
    $ brew install ngc224/txtmsk/txtmsk


Usage
-----
Set Password (Mac OS X Keychain)

    $ txtmsk -p

Encrypt

    $ txtmsk aaaaaaaaaaaaaaaaaaaa
    GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc

or Stdin

    $ cat ~/.aws/credentials | txtmsk > credentials.txtmsk

Decrypt

    $ txtmsk -d GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
    aaaaaaaaaaaaaaaaaaaa

or Stdin

    $ cat credentials.txtmsk | txtmsk > ~/.aws/credentials

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