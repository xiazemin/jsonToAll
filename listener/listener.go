package listener

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/parser"
	"strings"
)

type Node struct {
	Type string
	Value string
}

type Listener struct {
	*parser.BaseJSONListener
	gocodeMap map[antlr.Tree]Node `json:"gocode_map"`
	JsonStr string
	Target Target
}

func NewJsonToGoListener(t Target) *Listener {
	return &Listener{
		BaseJSONListener:&parser.BaseJSONListener{},
		gocodeMap:           make(map[antlr.Tree]Node),
		JsonStr :"",
		Target:t,
	}
}

func (l*Listener)PrintGocodeMap(){
	for k,v:=range l.gocodeMap{
		if false {
			fmt.Println(fmt.Sprintf("%T", k), v)
		}
	}
}

// VisitTerminal is called when a terminal node is visited.
func (l*Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (l*Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (l*Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (l*Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson is called when production json is entered.
func (l*Listener) EnterJson(ctx *parser.JsonContext) {}

// ExitJson is called when production json is exited.
func (l*Listener) ExitJson(ctx *parser.JsonContext) {
   l.JsonStr=l.Target.ExitJson(l.gocodeMap[ctx.Value()].Type,l.gocodeMap[ctx.Value()].Value)
}

// EnterObj is called when production obj is entered.
func (l*Listener) EnterObj(ctx *parser.ObjContext) {}

// ExitObj is called when production obj is exited.
func (l*Listener) ExitObj(ctx *parser.ObjContext) {
	sb := strings.Builder{}
	sb.WriteString(l.Target.PreExitObj("",""))
	for i,p:=range ctx.AllPair(){
		sb.WriteString(l.Target.ExitObj(l.gocodeMap[p].Type,l.gocodeMap[p].Value,i==len(ctx.AllPair())-1))
	}
	sb.WriteString(l.Target.PostExitObj("",""))
	l.gocodeMap[ctx]=Node{
		Type:  "struct",
		Value: sb.String(),
	}
}

// EnterPair is called when production pair is entered.
func (l*Listener) EnterPair(ctx *parser.PairContext) {}

// ExitPair is called when production pair is exited.
func (l*Listener) ExitPair(ctx *parser.PairContext) {
	l.gocodeMap[ctx]=Node{
		Type:  "KV",
		Value: l.Target.ExitPair(ctx.STRING().GetText(),l.gocodeMap[ctx.Value()].Type,l.gocodeMap[ctx.Value()].Value),
	}
	//fmt.Println(ctx.Value().GetText(),"======>",l.gocodeMap[ctx.Value()])
}

// EnterArr is called when production arr is entered.
func (l*Listener) EnterArr(ctx *parser.ArrContext) {}

// ExitArr is called when production arr is exited.
func (l*Listener) ExitArr(ctx *parser.ArrContext) {
	l.gocodeMap[ctx]=Node{
		Type:  "array",
		Value: l.Target.ExitArr(l.gocodeMap[ctx.Value(0)].Type ,l.gocodeMap[ctx.Value(0)].Value),
	}
	//fmt.Println(ctx.GetChild(0),ctx.Value(0))
}

// EnterValue is called when production value is entered.
func (l*Listener) EnterValue(ctx *parser.ValueContext) {}

// ExitValue is called when production value is exited.
func (l*Listener) ExitValue(ctx *parser.ValueContext) {
	if ctx.Arr()!=nil{
       l.gocodeMap[ctx]=l.gocodeMap[ctx.Arr()]
	}else if ctx.Obj()!=nil{
		l.gocodeMap[ctx]=l.gocodeMap[ctx.Obj()]
	}else if ctx.NUMBER()!=nil{
		l.gocodeMap[ctx]=Node{
			Type:  "float64",
			Value: l.Target.ExitValue("float64",ctx.NUMBER().GetText()) ,
		}
			//l.gocodeMap[ctx.NUMBER()]
	}else if ctx.STRING()!=nil{
		l.gocodeMap[ctx]=Node{
			Type:  "string",
			Value: l.Target.ExitValue("string",ctx.STRING().GetText()) ,
		}
			//l.gocodeMap[ctx.STRING()]
	}else{
		if ctx.GetText()=="true" || ctx.GetText()=="false"{
			l.gocodeMap[ctx]=Node{
				Type:  "bool",
				Value: l.Target.ExitValue("bool",ctx.GetText()) ,
			}
		}else if ctx.GetText()=="null"{
			l.gocodeMap[ctx]=Node{
				Type:  "null",
				Value: l.Target.ExitValue("null",ctx.GetText()) ,
			}
		}else{
			l.gocodeMap[ctx]=Node{
				Type:  "string",
				Value: l.Target.ExitValue("string",ctx.GetText()) ,
			}
		}
	}
}

