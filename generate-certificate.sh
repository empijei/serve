#!/bin/env bash
CONFPATH="$HOME/.servepj"
mkdir -p $CONFPATH
openssl req -x509 -newkey rsa:2048 -keyout "$CONFPATH/key.pem" -out "$CONFPATH/cert.pem" -days 3000 -nodes
