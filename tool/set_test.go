package tool

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T)  {
		slice1 := []string{"1", "2", "3", "6", "8"}
		slice2 := []string{"2", "3", "5", "0"}
		un := union(slice1, slice2)
		fmt.Println("slice1与slice2的并集为：", un)
		in := intersect(slice1, slice2)
		fmt.Println("slice1与slice2的交集为：", in)
		di := difference(slice1, slice2)
		fmt.Println("slice1与slice2的差集为：", di)
		di1 := difference(slice2, slice1)
		fmt.Println("slice2与slice1的差集为：", di1)
}
