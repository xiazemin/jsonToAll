java -jar ../calc/antlr-4.8-complete.jar -Dlanguage=Go -o parser JSON.g4
$go mod init github.com/xiazemin/json-parser/antlr/jsonToAll
go: creating new go.mod: module github.com/xiazemin/json-parser/antlr/jsonToAll


$go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/parser=./parser

$go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/listener=./listener

$go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/file=./file
$go mod edit -replace github.com/xiazemin/json-parser/antlr/jsonToAll/generator=./generator


export GOPROXY=https://goproxy.cn

go run ./