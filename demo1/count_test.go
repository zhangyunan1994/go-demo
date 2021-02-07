package demo1

import "testing"

func TestVar(t *testing.T)  {
	// -128 - 127
	var a int8 = -1
	var b int8 = -128 / a
	t.Log(b)
}

func TestConst(t *testing.T)  {
	// -128 - 127
	const a int8 = -1
	var b int8 = -128 / a
	t.Log(b)
}
