package demo1

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	dict := make(map[string]int)
	rangeMap(dict)

	dict["xx"] = 18
	rangeMap(dict)

	println(dict["xx"])

	dictSlice := map[string][]string{}
	rangeMapSlice(dictSlice)

	dictSlice["zyn"] = []string{"zhangyunan", "张瑀楠", "zhang.yunan"}
	rangeMapSlice(dictSlice)


	t.Log("success")
}

func rangeMapSlice(param map[string][]string)  {
	fmt.Println(param)
}

func rangeMap(map1 map[string]int)  {
	println("---------------------")
	for key, value := range map1 {
		println(key, value)
	}
}
