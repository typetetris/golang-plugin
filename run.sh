#!/bin/bash
GOPATH=$(pwd)
go build -buildmode=plugin testplugin
go install direct
bin/direct
go install indirect
bin/indirect
go build -buildmode=c-archive indirect_with_c
gcc -o indirect_with_c c_host_app.c indirect_with_c.a -ldl -lpthread
./indirect_with_c
