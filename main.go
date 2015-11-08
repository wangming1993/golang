package main

import (
	"fmt"
	"gostart/mystruct"
	"gostart/pointer"
	"time"
)

const EMAIL string = "mikewang@augmentum.com.cn"

func main() {
	fmt.Println("start go")
	variables()
	constants()
	forTest()
	ifElse()
	switchTest()
	arrayTest()
	slicesTest()
	mapTest()
	rangeTest()
	funcTest()
	closuresTest()
	recursiveTest()
	pointer.Pointer()
	mystruct.Mystruct()
}

func variables() {
	var a, b int = 1, 2
	fmt.Println(a, b)

	name := "mike"
	fmt.Println(name)
}

func constants() {
	fmt.Println("email:", EMAIL)
	const age int = 22
	fmt.Println("age:", age)
}

func forTest() {
	i := 1
	for i <= 3 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()
	for j := 1; j <= 3; j++ {
		fmt.Print(j, "\t")
	}
	fmt.Println()
	for {
		fmt.Println("use break to stop")
		break
	}
}

func ifElse() {
	if i := 10; i%2 == 0 {
		fmt.Println(i, "is odd")
	} else {
		fmt.Println(i, "is even")
	}
}

func switchTest() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("work day")
	}
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("afternoon")
	}
}

func arrayTest() {
	var a [5]int
	fmt.Println(a)
	a[2] = 3
	fmt.Println(a)

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	fmt.Println("len:", len(b))
}

func slicesTest() {
	slice := make([]string, 5)
	slice[0] = "I'm"
	slice[1] = "mike"
	fmt.Println(slice)
	slice = append(slice, "age")
	slice = append(slice, "22")
	fmt.Println(slice)
	fmt.Println("len:", len(slice))
}

func mapTest() {
	msg := make(map[string]string)
	msg["name"] = "mike"
	msg["age"] = "22"
	msg["gender"] = "male"
	fmt.Println(msg)
	fmt.Println(len(msg))
	delete(msg, "name")
	fmt.Println(len(msg))
}

func rangeTest() {
	msg := map[string]string{"name": "mike", "age": "22", "gender": "male"}
	for key, value := range msg {
		fmt.Println(key, "=>", value)
	}
}

func getMsg() (string, string, string) {
	msg := map[string]string{"name": "mike", "age": "22", "gender": "male"}
	return msg["name"], msg["age"], msg["gender"]
}

func printMsg(keys ...string) {
	msg := map[string]string{"name": "mike", "age": "22", "gender": "male"}
	for _, key := range keys {
		fmt.Println(msg[key])
	}
}
func funcTest() {
	name, age, male := getMsg()
	fmt.Println(name, age, male)
	printMsg("name")
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func closuresTest() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	next := intSeq()
	fmt.Println(next())
	fmt.Println(next())
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func recursiveTest() {
	n := 5
	fmt.Println(fact(n))
}
