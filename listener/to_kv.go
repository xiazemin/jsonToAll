package listener

type KVTarget struct {

}

// func (kv*KVTarget)ExitJson is called when production json is exited.
func (kv*KVTarget)ExitJson(typeStr,valStr string)string{ return typeStr+":"+valStr}
// func (kv*KVTarget)ExitObj is called when production obj is exited.
func (kv*KVTarget)PreExitObj(typeStr,valStr string)string{ return typeStr+":"+valStr}
func (kv*KVTarget)ExitObj(typeStr,valStr string,isEnd bool,bIsMap bool)string{ return typeStr+":"+valStr}
func (kv*KVTarget)PostExitObj(typeStr,valStr string)string{ return typeStr+":"+valStr}
// func (kv*KVTarget)ExitPair is called when production pair is exited.
func (kv*KVTarget)ExitPair(index int,keyStr,typeStr,valStr,valType string)(string,string){ return "",keyStr+"#"+typeStr+":"+valStr}
// func (kv*KVTarget)ExitArr is called when production arr is exited.
func (kv*KVTarget)ExitArr(typeStr,valStr string)string{ return typeStr+":"+valStr}
// func (kv*KVTarget)ExitValue is called when production value is exited.
func (kv*KVTarget)ExitValue(typeStr,valStr string)string{ return typeStr+":"+valStr}
