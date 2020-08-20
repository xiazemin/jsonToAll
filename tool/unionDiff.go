package tool

func UnionDiffJson(jsonStr1, jsonStr2 string) string  {
	diff1:=DiffJSON(jsonStr1,jsonStr2)
	diff2:=DiffJSON(jsonStr2,jsonStr1)
	return UnionJSON(diff1,diff2)
}
