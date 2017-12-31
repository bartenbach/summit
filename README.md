# summit
summit is a command line checksum validator.
![summmit](summit.png "Picture of summit in action")

Ever try and validate file integrity checksums before by just looking at them?  Yeah...that seems foolproof.  Well, now you don't have to anymore.  You can do it in an automated fashion!

## usage
![help](help.png "Help menu")
The help menu should be sufficient enough, but you will generally use 3 arguments.

`-t` this is the type of hash to use to validate your checksum.  Typically md5 or sha1.
`-f` this is the path to the file you want to verify.
`-e` this is the *correct* checksum you are hoping that your file has, if all went well.

## examples
`summit -t md5 -f ./archlinux.iso -e d0ae8c4b4a037238e5f0880243d4619e`
`summit -t sha1 -f ~/.bashrc`
The `-e` flag is optional in a sense.  You can still get the hash of a file without specifying it.
