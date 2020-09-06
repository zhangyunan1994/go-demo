package demo1

import (
	"testing"
)

func TestArray(t *testing.T) {
	var array [5]int
	RangeIntArray(&array)

	array = [5]int{1,2,3,4,5}
	RangeIntArray(&array)

	array2 := [...]int{6,7,8,9,10}
	RangeIntArray(&array2)

	array3 := [5]int{2:23}
	RangeIntArray(&array3)

}

func RangeIntArray(array *[5]int)  {
	println("==========================")
	for index, item := range *array {
		println(index, item)
	}
}

