#!/usr/bin/env bash

# Go get the cfssl toolkit
if [ -z "$GOBIN" ] || [ -z "$GOPATH" ]; then
	echo "You dont seem to have a correctly configured Go toolchain"
	exit 1
fi
if ! hash cfssl 2>/dev/null; then
	echo "installing cfssl to GOBIN"
 	go get -u github.com/cloudflare/cfssl/cmd/...
fi
cd .cfssl
cfssl genkey -initca csr.json | cfssljson -bare ca
cfssl gencert -ca ca.pem -ca-key ca-key.pem -hostname=encryption-server.yoti-test,0.0.0.0 csr.json | cfssljson -bare server
cfssl gencert -ca ca.pem -ca-key ca-key.pem -hostname=encryption-client.yoti-test,0.0.0.0 csr.json | cfssljson -bare client

echo "Done"
exit 0