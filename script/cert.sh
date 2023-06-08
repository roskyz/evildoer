#!/bin/bash

script_path="$(cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P)"
working_directory="$script_path/../conf/cert"
cd $working_directory
working_directory=`pwd -P`
echo -e "Change working directory to $working_directory \nBegin to issue self-signed certificates..."

# CA
echo "Generate CA's private key and self-signed certificate"
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CN/ST=HongKong/L=HongKong/O=SelfProject/OU=Business/CN=localhost.hk/emailAddress=none@localhost.hk"

echo "Output CA's self-signed certificate"
openssl x509 -in ca-cert.pem -nroot -text

# Server
echo "Generate server's private key and certificate signing request (CSR)"
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=CN/ST=HongKong/L=HongKong/O=Good/OU=Education/CN=*.web.local/emailAddress=admin@web.local"

echo "Use CA's private key to sign server' CSR and get back the signed certificate"
openssl x509 -req -extfile <(printf "subjectAltName=DNS:*.web.local") -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

echo "Output server's signed certificate"
openssl x509 -in server-cert.pem -nroot -text

# Client
echo "Generate client's private key and certificate signing request (CSR)"
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=CN/ST=HongKong/L=HongKong/O=Good/OU=Education/CN=*.web.local/emailAddress=admin@web.local"

echo "Use CA's private key to sign client' CSR and get back the signed certificate"
openssl x509 -req -extfile <(printf "subjectAltName=DNS:*.web.local") -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem

echo "Output client's signed certificate"
openssl x509 -in client-cert.pem -nroot -text