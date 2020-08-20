package tool

import (
	"fmt"
	"testing"
)

func TestMergeJSON(t *testing.T) {
	fmt.Println(MergeJSON(
		"[{\"1\":1},{\"2\":2}]",
		"[{\"1\":1},{\"2\":2},{\"3\":3},{\"4\":[{\"a\":3}]}]"))

	//坑 map 和list 是反的,list以第二个为准,map以第一个为准

	fmt.Println(MergeJSON(
		"[{\"1\":1},{\"2\":2},{\"2\":4,\"4\":4,\"3\":3},{\"4\":[{\"a\":3}]}]",
		"[{\"1\":1},{\"3\":4},{\"2\":4,\"4\":4,\"3\":4}]"))

}
