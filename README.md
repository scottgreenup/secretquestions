# Secret Questions

This is for generating secret answers for secret questions. I usually make the
secret question custom, using something like "What is the answer?" and then
fill in the answer with the output of this program. I also store the answer in
a password manager, or offline in a book, etc...

The only requirement is that you source a wordlist from somewhere and have it
at `./wordlist.txt`

## Getting a Word List

Using aspell:
```
aspell dump master > wordlist.txt
```

Download one
 - https://github.com/berzerk0/Probable-Wordlists/tree/master/Dictionary-Style
 - https://weakpass.com/wordlist/1239
 - https://github.com/dwyl/english-words

You could check `/usr/share/dict` to see if you already have one.
```
ls /usr/share/dict
ln -sT /usr/share/dict/X wordlist.txt
```
