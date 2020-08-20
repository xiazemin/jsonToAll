module github.com/xiazemin/json-parser/antlr/jsonToAll/listener

go 1.13

require (
	github.com/antlr/antlr4 v0.0.0-20200819162015-533f63382ac9
	github.com/xiazemin/json-parser/antlr/jsonToAll/parser v0.0.0-00010101000000-000000000000
)

replace github.com/xiazemin/json-parser/antlr/jsonToAll/parser => ../parser
