package main

import (
	"fmt"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/file"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/generator"

	"log"
)

func main()  {
	strOri,strGen:=generator.Gen("./t.json","go")
	fmt.Println(strOri)
	log.Println(strGen)

	file.PutGoLang("./gen/t.go",strGen)
	strJsonOri,strJsonGen:=generator.Gen("./t.json","json")
	fmt.Println(strJsonGen)
	log.Println(strJsonOri)

	file.PutJson("./gen/t.json",strJsonGen)
}

