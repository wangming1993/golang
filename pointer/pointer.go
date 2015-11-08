package pointer

import (
	"fmt"
)

func Pointer() {
	fmt.Println("go pointer")
	n := 1
	fmt.Println(n)
	setValue(n)
	fmt.Println(n)
	setValueUserPointer(&n)
	fmt.Println(n)
	fmt.Println(&n)
}

func setValue(n int) {
	n = 0
}

func setValueUserPointer(n *int) {
	*n = 0
}
