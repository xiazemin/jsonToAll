package generator

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/file"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/listener"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/parser"
)

func Gen(name ,genType string)(string,string,[]string){
	s:=string(file.GetJson(name))
	in:=antlr.NewInputStream(
		s,
	)
	lex:=parser.NewJSONLexer(in)
	stream:=antlr.NewCommonTokenStream(lex,antlr.TokenDefaultChannel)
	par:=parser.NewJSONParser(stream)
	var l *listener.Listener
	switch genType {
	case "json":
		l=listener.NewJsonToGoListener(&listener.JsonTarget{})
	case "go":
		l=listener.NewJsonToGoListener(&listener.GoTarget{})
	case "kv":
		l=listener.NewJsonToGoListener(&listener.KVTarget{})
	case "idl":
		l=listener.NewJsonToGoListener(&listener.IdlTarget{})

	}

	antlr.ParseTreeWalkerDefault.Walk(l,par.Json())
	//l.PrintGocodeMap()
	return s,l.JsonStr,l.SubStructs
}
