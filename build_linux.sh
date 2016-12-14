#!/bin/bash

go build --ldflags '-extldflags "-static"' mailgunme.go
mv mailgunme mgm.linux
