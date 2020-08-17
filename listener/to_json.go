package listener

type JsonTarget struct {

}

// func (j*JsonTarget)ExitJson is called when production json is exited.
func (j*JsonTarget)ExitJson(typeStr,valStr string)string{
	/*if typeStr=="struct"{
      return "{"+valStr+"}"
	}else if typeStr=="array"{
		return "["+valStr+"]"
	}
	return typeStr+":"+valStr
	*/
	return valStr
}
// func (j*JsonTarget)ExitObj is called when production obj is exited.
func (j*JsonTarget)PreExitObj(typeStr,valStr string)string{
	//return "{"+valStr+"}"
	return "{\n"
}
func (j*JsonTarget)ExitObj(typeStr,valStr string,isEnd bool)string{
	//return "{"+valStr+"}"
	//fmt.Println(typeStr,valStr,isEnd)
	if isEnd{
		return valStr+"\n"
	}
	return valStr+",\n"
}
func (j*JsonTarget)PostExitObj(typeStr,valStr string)string{
	//return "{"+valStr+"}"
	return "}"
}
// func (j*JsonTarget)ExitPair is called when production pair is exited.
func (j*JsonTarget)ExitPair(index int,keyStr,typeStr,valStr,valType string)(string,string){
	return "",keyStr+":"+valStr
}
// func (j*JsonTarget)ExitArr is called when production arr is exited.
func (j*JsonTarget)ExitArr(typeStr,valStr string)string{
	return "[\n"+valStr+"\n]"
}
// func (j*JsonTarget)ExitValue is called when production value is exited.
func (j*JsonTarget)ExitValue(typeStr,valStr string)string{
	return valStr
}