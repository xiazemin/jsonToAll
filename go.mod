module github.com/xiazemin/json-parser/antlr/jsonToAll

go 1.13

replace github.com/xiazemin/json-parser/antlr/jsonToAll/parser => ./parser

replace github.com/xiazemin/json-parser/antlr/jsonToAll/listener => ./listener

replace github.com/xiazemin/json-parser/antlr/jsonToAll/file => ./file

require (
	github.com/antlr/antlr4 v0.0.0-20200819162015-533f63382ac9
	github.com/xiazemin/json-parser/antlr/jsonToAll/file v0.0.0-00010101000000-000000000000
	github.com/xiazemin/json-parser/antlr/jsonToAll/generator v0.0.0-00010101000000-000000000000
	github.com/xiazemin/json-parser/antlr/jsonToAll/listener v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/xiazemin/json-parser/antlr/jsonToAll/generator => ./generator
