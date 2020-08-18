module github.com/xiazemin/json-parser/antlr/jsonToAll/cmd

go 1.13

replace github.com/xiazemin/json-parser/antlr/jsonToAll/cmd/cmd => ./cmd

require (
	github.com/antlr/antlr4 v0.0.0-20200801005519-2ba38605b949 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/xiazemin/json-parser/antlr/jsonToAll/file v0.0.0-00010101000000-000000000000
	github.com/xiazemin/json-parser/antlr/jsonToAll/generator v0.0.0-00010101000000-000000000000
	github.com/xiazemin/json-parser/antlr/jsonToAll/listener v0.0.0-00010101000000-000000000000 // indirect
	github.com/xiazemin/json-parser/antlr/jsonToAll/parser v0.0.0-00010101000000-000000000000 // indirect
	github.com/xiazemin/json-parser/antlr/jsonToAll/tool v0.0.0-00010101000000-000000000000
)

replace github.com/xiazemin/json-parser/antlr/jsonToAll/file => ../file

replace github.com/xiazemin/json-parser/antlr/jsonToAll/generator => ../generator

replace github.com/xiazemin/json-parser/antlr/jsonToAll/listener => ../listener

replace github.com/xiazemin/json-parser/antlr/jsonToAll/parser => ../parser

replace github.com/xiazemin/json-parser/antlr/jsonToAll/tool => ../tool
