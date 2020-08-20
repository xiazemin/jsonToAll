package tool

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
)

//MergeJSON 合并两个json，用 jsonStr1 的值覆盖 jsonStr2 的值，并舍弃jsonStr2中不存在的内容。注意剔除json前后的空字符
func MergeJSON(jsonStr1, jsonStr2 string) string {
	//先看看两个json类型一样不，根据第一个字符判断
	if jsonStr1[:1] != jsonStr2[:1] {
		return jsonStr2 //合并失败直接返回第二个json
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

			result := mergeSorttedMap(obj1, obj2) //调用递归合并方法
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

			result := mergeMap(obj1, obj2) //调用递归合并方法
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
		result := mergeList(obj1, obj2) //调用递归合并方法
		//将结果转回字符串返回
		jsonBytes, err := json.Marshal(result)
		checkErr(err)
		return string(jsonBytes)
	}
	return jsonStr2 //合并失败直接返回第二个json
}

//mergeMap map递归合并操作
func mergeSorttedMap(map1, map2 *orderedmap.OrderedMap) *orderedmap.OrderedMap {
	resultMap := orderedmap.New()    //先创建空结果
	for _, k2 := range map2.Keys() { //提取map2的键值
		v2, exist := map2.Get(k2)
		if !exist {
			continue
		}
		resultMap.Set(k2, v2)     //先按map2填充结果
		v1, exist := map1.Get(k2) //提取map1的对应内容
		if !exist {               //不存在则跳过
			continue
		}
		switch v2.(type) { //检查类型
		case string, float64, bool: //基本类型直接用map1的内容为结果赋值
			resultMap.Set(k2, v1)
		case []interface{}: //list类型调用list的递归合并操作，并赋值给结果
			resultMap.Set(k2, mergeList(v1.([]interface{}), v2.([]interface{})))
		case map[string]interface{}: //map类型调用map递归合并操作，并赋值给结果
			fmt.Println("\033[31;1m fatal error :map[string]interface{} \033[0m") //("\\e[31;1m fatal error \\e[0m\n")
			//resultMap[k2] = mergeSorttedMap(v1.(map[string]interface{}), v2.(map[string]interface{}))
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
			resultMap.Set(k2, mergeSorttedMap(&v11, &v22))
		}
	}
	return resultMap //返回最终合并结果对象
}

//mergeMap map递归合并操作,map1 新值，map2原始值
func mergeMap(map1, map2 map[string]interface{}) map[string]interface{} {
	resultMap := make(map[string]interface{}) //先创建空结果
	//针对置空的需求
	if len(map1)==0{
		fmt.Println("\033[31;1m xiazemin merge map delete data \033[0m")
		return resultMap;
	}
	for k2, v2 := range map2 {                //提取map2的键值
		resultMap[k2] = v2    //先按map2填充结果
		v1, exist := map1[k2] //提取map1的对应内容
		if !exist {           //不存在则跳过
			continue
		}
		switch v2.(type) { //检查类型
		case string, float64, bool: //基本类型直接用map1的内容为结果赋值
			resultMap[k2] = v1
		case []interface{}: //list类型调用list的递归合并操作，并赋值给结果
			fmt.Println("\033[31;1m fatal error \033[0m") //("\\e[31;1m fatal error \\e[0m\n")
			resultMap[k2] = mergeList(v1.([]interface{}), v2.([]interface{}))
		case map[string]interface{}: //map类型调用map递归合并操作，并赋值给结果
			resultMap[k2] = mergeMap(v1.(map[string]interface{}), v2.(map[string]interface{}))
		}
	}
	return resultMap //返回最终合并结果对象
}

//mergeList list递归合并操作
func mergeList(list1, list2 []interface{}) []interface{} {
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
			resultList[index] = list1[index]
		case []interface{}: //list类型调用list的递归合并操作，并赋值给结果
			resultList[index] = mergeList(list1[index].([]interface{}), item2.([]interface{}))
		case map[string]interface{}: //map类型调用map递归合并操作，并赋值给结果
			if SortedMap {
				fmt.Println("\033[31;1m list type fatal error initEnv.SortedMap \033[0m")
			}
			resultList[index] = mergeMap(list1[index].(map[string]interface{}), item2.(map[string]interface{}))
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
			resultList[index] = mergeSorttedMap(&v11, &v22)
		}
	}
	fmt.Println(index, len(list1))
	if index == len(list1) { //list2长度大于list1直接merge
		resultList = append(resultList, list2[index:]...)
	}
	return resultList
}

func checkErr(e error) {
	fmt.Println(e)
}
