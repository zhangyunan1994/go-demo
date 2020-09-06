package demo1

import "testing"

func TestSlice(t *testing.T) {
	slice1 := make([]string, 2)
	RangeStringSlice(&slice1)

	slice2 := []int{}
	RangeIntSlice(&slice2)

	slice3 := []int{1,2,3,4,5}
	RangeIntSlice(&slice3)

	slice4 := slice3[1:3]
	RangeIntSlice(&slice4)

	slice4 = append(slice4, 8)
	slice4 = append(slice4, 9)
	slice4 = append(slice4, 10)
	slice4 = append(slice4, 11)
	RangeIntSlice(&slice4)
	RangeIntSlice(&slice3)

	slice3 = append(slice3, 1000)
	RangeIntSlice(&slice3)
	RangeIntSlice(&slice4)


	t.Log("success")
}

func RangeIntSlice(slice *[]int)  {
	println("------------------")
	for index, value := range *slice {
		println(index, value)
	}
}

func RangeStringSlice(slice *[]string)  {
	println("------------------")
	for index, value := range *slice {
		println(index, value)
	}
}
