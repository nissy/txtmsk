# txtmsk

Usage
-----
Set Password

    $ txtmsk -p

Encrypt

    $ echo "aaaaaaaaaaaaaaaaaaaa" | txtmsk
    GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc
    
Decrypt

    $ echo "GFKm0AD9g0yyUdCc6cq44sX+D6CAyWnqzoxa4jU0rZdC4ZOc" | txtmsk -d
    aaaaaaaaaaaaaaaaaaaa
    