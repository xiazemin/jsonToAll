package tool

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
)

//DiffJSON  jsonStr1-jsonStr2
func DiffJSON(jsonStr1, jsonStr2 string) string {
	//先看看两个json类型一样不，根据第一个字符判断
	if jsonStr1[:1] != jsonStr2[:1] {
		return jsonStr1 //合并
	}
	var err error
	if jsonStr2[:1] == "{" { //map类型
		fmt.Println(SortedMap)
		if SortedMap {
			obj1 := orderedmap.New()
			obj2 := orderedmap.New()
			err = json.Unmarshal([]byte(jsonStr1), &obj1)
			checkErr(err)
			err = json.Unmarshal([]byte(jsonStr2), &obj2)
			checkErr(err)

			result := diffSorttedMap(obj1, obj2) //调用递归合并方法
			//将结果转回字符串返回
			jsonBytes, err := json.Marshal(result)
			checkErr(err)
			return string(jsonBytes)
		} else {
			obj1 := make(map[string]interface{})
			obj2 := make(map[string]interface{})
			err = json.Unmarshal([]byte(jsonStr1), &obj1)
			checkErr(err)
			err = json.Unmarshal([]byte(jsonStr2), &obj2)
			checkErr(err)

			result := diffMap(obj1, obj2) //调用递归合并方法
			//将结果转回字符串返回
			jsonBytes, err := json.Marshal(result)
			checkErr(err)
			return string(jsonBytes)
		}
	}
	if jsonStr2[:1] == "[" { //list类型
		obj1 := make([]interface{}, 0)
		obj2 := make([]interface{}, 0)
		err = json.Unmarshal([]byte(jsonStr1), &obj1)
		checkErr(err)
		err = json.Unmarshal([]byte(jsonStr2), &obj2)
		checkErr(err)

		fmt.Println(obj1, obj2)
		result := diffList(obj1, obj2) //调用递归合并方法
		//将结果转回字符串返回
		jsonBytes, err := json.Marshal(result)
		checkErr(err)
		return string(jsonBytes)
	}
	return jsonStr2 //合并失败直接返回第二个json
}

//diffMap map递归合并操作
func diffSorttedMap(map1, map2 *orderedmap.OrderedMap) *orderedmap.OrderedMap {
	resultMap := orderedmap.New()    //先创建空结果
	allKeys:=union(map1.Keys(),map2.Keys())
	for _,k:=range allKeys{
		v2, exist2 := map2.Get(k) //提取map2的键值
		v1, exist1 := map1.Get(k) //提取map1的对应内容
		if exist1 && !exist2 {
			resultMap.Set(k, v1)
		}else if !exist1 && exist2{
			//resultMap.Set(k,v1)
		}else {
			switch v2.(type) { //检查类型
			case string, float64, bool: //基本类型直接用map1的内容为结果赋值
				//resultMap.Set(k, v1)
			case []interface{}: //list类型调用list的递归合并操作，并赋值给结果
				resultMap.Set(k, diffList(v1.([]interface{}), v2.([]interface{})))
			case map[string]interface{}: //map类型调用map递归合并操作，并赋值给结果
				fmt.Println("\033[31;1m fatal error :map[string]interface{} \033[0m") //("\\e[31;1m fatal error \\e[0m\n")
				//resultMap[k2] = diffSorttedMap(v1.(map[string]interface{}), v2.(map[string]interface{}))
			case orderedmap.OrderedMap:
				v11, ok := v1.(orderedmap.OrderedMap)
				if !ok {
					fmt.Println("\033[31;1m v1 type fatal error: orderedmap.OrderedMap  v11 \033[0m")
					continue
				}
				v22, ok := v2.(orderedmap.OrderedMap)
				if !ok {
					fmt.Println("\033[31;1m v2 type fatal error orderedmap.OrderedMap v22 \033[0m")
					continue
				}
				resultMap.Set(k, diffSorttedMap(&v11, &v22))
			}
		}
	}
	return resultMap //返回最终合并结果对象
}

//diffMap map递归合并操作,map1 新值，map2原始值
func diffMap(map1, map2 map[string]interface{}) map[string]interface{} {
	resultMap := make(map[string]interface{}) //先创建空结果

	allKeys:=difference(getKeys(map1),getKeys(map2))
	for _,k:=range allKeys{
		v1, exist1 := map1[k] //提取map1的对应内容
		v2, exist2 := map2[k] //提取map2的对应内容

		if exist1 && !exist2{
			resultMap[k]=v1
		}else if !exist1 && exist2{
			//resultMap[k]=v2
		}else{
			switch v2.(type) { //检查类型
			case string, float64, bool: //基本类型直接用map1的内容为结果赋值
				//resultMap[k] = v1
			case []interface{}: //list类型调用list的递归合并操作，并赋值给结果
				fmt.Println("\033[31;1m fatal error \033[0m") //("\\e[31;1m fatal error \\e[0m\n")
				resultMap[k] = diffList(v1.([]interface{}), v2.([]interface{}))
			case map[string]interface{}: //map类型调用map递归合并操作，并赋值给结果
				resultMap[k] = diffMap(v1.(map[string]interface{}), v2.(map[string]interface{}))
			}
		}
	}
	return resultMap //返回最终合并结果对象
}

//diffList list递归合并操作,list 还是 合并操作，不然diff不全
func diffList(list1, list2 []interface{}) []interface{} {
	resultList := make([]interface{}, 0) //先创建空结果
	var index int
	var item2 interface{}
	for index, item2 = range list2 { //提取list2的索引以及内容
		if index == len(list1) { //若list2内容过长，则就此截断，以list1为准
			break
		}
		resultList = append(resultList, item2) //先按list2填充结果
		switch item2.(type) {                  //检查类型
		case string, float64, bool: //基本类型直接用list1的内容为结果赋值
			//resultList[index] = list1[index]
		case []interface{}: //list类型调用list的递归合并操作，并赋值给结果
			resultList[index] = diffList(list1[index].([]interface{}), item2.([]interface{}))
		case map[string]interface{}: //map类型调用map递归合并操作，并赋值给结果
			if SortedMap {
				fmt.Println("\033[31;1m list type fatal error initEnv.SortedMap \033[0m")
			}
			resultList[index] = diffMap(list1[index].(map[string]interface{}), item2.(map[string]interface{}))
		case orderedmap.OrderedMap:
			v11, ok := list1[index].(orderedmap.OrderedMap)
			if !ok {
				fmt.Println("\033[31;1m v1 type fatal error orderedmap.OrderedMap \033[0m")
				continue
			}
			v22, ok := item2.(orderedmap.OrderedMap)
			if !ok {
				fmt.Println("\033[31;1m v2 type fatal error  orderedmap.OrderedMap \033[0m")
				continue
			}
			resultList[index] = diffSorttedMap(&v11, &v22)
		}
	}
	fmt.Println(index, len(list1))
	if len(list2)> len(list1) { //list2长度大于list1直接merge
		//resultList = append(resultList, list2[index:]...)
	}else if len(list1)>len(list2){
		resultList = append(resultList, list1[index+1:]...)
	}


	return resultList
}
