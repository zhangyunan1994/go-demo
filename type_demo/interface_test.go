package type_demo

import (
	"fmt"
	"testing"
)

type notifier interface {
	notify()
}

type userNotify struct {
	Name string
}

func (u userNotify) notify ()  {
	fmt.Println("receiver ", u.Name)
}

func sendNotify(se notifier)  {
	se.notify()
}

func TestI(t *testing.T)  {

	u := userNotify{"zn"}
	sendNotify(u)


	t.Log("success")
}
