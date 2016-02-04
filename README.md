txtmsk
=======
txtmsk encrypts the text.

Mac OS X only !!!


Install
-----
    $ brew install ngc224/txtmsk/txtmsk


Usage
-----
Set Password (Mac OS X Keychain)

    $ txtmsk -p

Encrypt

    $ echo "aaaaaaaaaaaaaaaaaaaa" | txtmsk
    GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
    
Decrypt

    $ echo "GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc" | txtmsk -d
    aaaaaaaaaaaaaaaaaaaa
    