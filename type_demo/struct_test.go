package type_demo

import (
	"fmt"
	"testing"
)

type user struct {
	Name string
	Age int
	Sex int
}

func (u user) toString() {
	fmt.Println(u)
}

func (u user) changeAge(age int) {
	u.Age = age
}

func (u *user) changeName(name string)  {
	u.Name = name
}

type admin struct {
	user
	level string
}


func TestStruct(t *testing.T)  {
	zyn := user{"张瑀楠", 18, 1}
	zyn.toString()
	zyn.changeAge(26)
	zyn.toString()
	zyn.changeName("zzzkjjsk")
	zyn.toString()

	ll := user{Name: "zz", Age: 15, Sex: 56}
	fmt.Println(ll)

	admin1 := admin{user{Name: "zyn"}, "admin"}

	fmt.Println(admin1)

	t.Log("success")
}