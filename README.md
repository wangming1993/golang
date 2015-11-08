# golang

## go 包管理
> go中的规定，　源码写在src目录下面

### 代码结构：
    src
        gostart
            mystruct
                mystruct.go
            pointer
                pointer.go
            main.go
            Makefile
     
```golang
#main.go
package main
import (
    "fmt"
    "gostart/mystruct"
    "gostart/pointer"
    "time"
)
func main() {
    fmt.Println("start go")
    pointer.Pointer()
    mystruct.Mystruct()
}
```
```golang
#mystruct.go
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
```

```golang
#pointer.go
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
```

**特别需要注意的是，go中并没有public private等, 而是以首字母的大小写来区分, 如果需要在包外使用, 则需要将它的首字母大写以声明为public**

**go支持匿名函数,闭包和递归调用**
**分片比数组更为常用**