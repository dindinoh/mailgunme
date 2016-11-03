#!/bin/bash

go get github.com/mailgun/mailgun-go github.com/mitchellh/go-homedir gopkg.in/gcfg.v1

GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build --ldflags '-extldflags "-static"' mailgunme.go
