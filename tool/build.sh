#!/bin/bash
$ go mod init github.com/xiazemin/json-parser/antlr/jsonToAll/tool
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/jsonToAll/tool


$ cobra add union
union created at /Users/didi/PhpstormProjects/c/json-parser/antlr/jsonToAll/cmdlocalhost:cmd
didi$
$ cobra add diff
diff created at /Users/didi/PhpstormProjects/c/json-parser/antlr/jsonToAll/cmdlocalhost:cmd
didi$
$ cobra add uniondiff
uniondiff created at /Users/didi/PhpstormProjects/c/json-parser/antlr/jsonToAll/cmdlocalhost:cmd
didi$
$ go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/tool=../tool