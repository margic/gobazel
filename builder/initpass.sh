#!/bin/bash

echo Init Pass Store

# create pgp config file
umask 0277
cat << EOF > /tmp/gpg-genkey.conf
%echo Generating a package signing key
%no-protection
%no-ask-passphrase
Key-Type: DSA
Key-Length: 1024
Subkey-Type: ELG-E
Subkey-Length: 2048
Name-Real:  `hostname --fqdn`
Name-Email: builder@`hostname --fqdn`
Expire-Date: 0
%commit
%echo Done
EOF
umask 0002

gpg2 --batch --output gpg-builder.rev --full-gen-key /tmp/gpg-genkey.conf
rm /tmp/gpg-genkey.conf

pass init builder@$(hostname --fqdn)

echo Done Init Pass Store