package main

import (
	"fmt"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/file"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/generator"

	"log"
)

func main()  {
	strOri,strGen,_:=generator.Gen("./t.json","go")
	fmt.Println(strOri)
	log.Println(strGen)
	file.PutGoLang("./gen/t.go",strGen)

	strJsonOri,strJsonGen,_:=generator.Gen("./t.json","json")
	fmt.Println(strJsonGen)
	log.Println(strJsonOri)
	file.PutJson("./gen/t.json",strJsonGen)


	strIdlOri,strIdlGen,subStructs:=generator.Gen("./t.json","idl")
	fmt.Println(strIdlOri)
	log.Println(strIdlGen)
	//fmt.Println(subStructs,len(subStructs))
	file.PutIdl("./gen/t.idl",subStructs,strIdlGen)
}

