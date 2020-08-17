java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser JSON.g4
$go mod init github.com/xiazemin/json-parser/antlr/jsonToAll
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/jsonToAll


$go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/parser=./parser

$go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/listener=./listener

$go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/file=./file
$go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/generator=./generator


export GOPROXY=https://goproxy.cn

go run ./


https://walmartlabs.github.io/json-to-simple-graphql-schema/

$ cobra init .  --pkg-name  github.com/xiazemin/json-parser/antlr/jsonToAll/cmd
Your Cobra applicaton is ready at
/Users/didi/PhpstormProjects/c/json-parser/antlr/jsonToAll/cmd

$ go mod init github.com/xiazemin/json-parser/antlr/jsonToAll/cmd
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/jsonToAll/cmd


$ go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/cmd/cmd=./cmd
export GOPROXY=https://goproxy.cn

go run ./

$ cobra add toIdl
toIdl created at /Users/didi/PhpstormProjects/c/json-parser/antlr/jsonToAll/cmd

 go build -o jsonToIdl main.go

https://o-my-chenjian.com/2017/09/20/Using-Cobra-With-Golang/

$ go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/file=../file
$ go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/generator=../generator

$ go build -o jsonToIdl main.go

 ./jsonToIdl toIdl -j ../t.json
  ./jsonToIdl toIdl -j ../t.json -d ./cmd