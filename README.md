**Usage**

Stack type is very simple to use. This is a very basic example:

```
import "gitlab.com/matteo.angelino/stack"
import "fmt"

func main(){
    s := new(stack.Stack)
    s.Push(10)
    s.Push(15)
    s.Push(20)
    fmt.Println("Stack size : ",s.Size())
    e := s.Pop()
    fmt.Println("Stack size : ",s.Size())
    fmt.Println("Element : ",e.Value())
}
```