![github-banner](https://user-images.githubusercontent.com/96031819/218027133-7bfd8533-bb05-482b-87b1-7482d95b8fe7.png)

# curvy
`ed25519` -> `x25519` 

## what?
`ed25519` is for signing, `x25519` is for encrypting. 

this thing generates a set of `ed25519` keys, then uses them to derive a set of `x25519` keys.

## what? -vv
ok so if you generate some ssh keys right now, you're going to get some type of dollar store brand rsa keys, unless you tell `ssh-keygen` that you want to make a set of the good ones.  you can use these to 'sign' things, and your friends can use your 'signature' to verify for themselves that nothing was tampered with.

`x25519` is for encrypting, and like `ed25519` there is a public key and a private key. `x25519` is for encrypting, which is different than signing.

## What? -vv -still

eventually I'd like to take this a bit further and turn this repo into an `edx25519` tool like Fellipo described in this [article](https://words.filippo.io/using-ed25519-keys-for-encryption/)

i guess it's not a good idea to reuse keys and such, so definitely don't use this, and if you get jammed up using this, just tell them you were holding it for a friend.
