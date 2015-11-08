package mystruct

import (
	"fmt"
)

type msg struct {
	name   string
	age    int
	gender string
}

func Mystruct() {
	fmt.Println("go mystruct")
	myMsg := msg{name: "mike", age: 22, gender: "male"}
	fmt.Println(myMsg.name)
}
