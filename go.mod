module github.com/xiazemin/json-parser/antlr/jsonToAll

go 1.13

replace github.com/xiazemin/json-parser/antlr/jsonToAll/parser => ./parser

replace github.com/xiazemin/json-parser/antlr/jsonToAll/listener => ./listener

replace github.com/xiazemin/json-parser/antlr/jsonToAll/file => ./file

require github.com/antlr/antlr4 v0.0.0-20200712162734-eb1adaa8a7a6

replace github.com/xiazemin/json-parser/antlr/jsonToAll/generator => ./generator
