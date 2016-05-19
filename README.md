**Usage**

Stack type is very simple to use. This is a very basic example:

```
import "github.com/matteo.angelino/stack"
import "fmt"

func main(){
    var s stack.Stack
    s.Push(10)
    s.Push(15)
    s.Push(20)
    fmt.Println("Stack size : ",s.Size())
    e := s.Pop()
    fmt.Println("Stack size : ",s.Size())
    fmt.Println("Element : ",e.Value())
}
```

**Methods**

On type Stack are defined the following methods:

```
Push(interface{}) :  Add element at top of the stack
Pop() *Element    :  Remove the first (top) element from the stack
Top() *Element    :  Get the top element from the stack (without remove it)
Size() int        :  Return the size of the stack
```

For compatibility with the go type *container.list* are also implemented the two following methods:

```
Front() *Element : alias of Pop
Next() *Element  : alias of Pop
```

**Access element value**

Stack type are composed by type *Element linked together. To access the value stored in the element use the *Value()* method. *Value()* method return type *interface{}* so you need to convert that type to the correct value type. For example:

```
    var s stack.Stack
    s.Push(10)
    s.Push("This is a string")
    
    for e := s.Pop(); e != nil; e = s.Pop(){
        v, ok := e.Value().(int)
        if ok {
            fmt.Println("Type is int : ",v)
        }
        v, ok = e.Value().(string)
        if ok {
            fmt.Println("Type is string : ",v)   
        }
    } 
```

**Testing**

In order tu run test suite, execute the following command:

```
go test -v gitlab.com/matteo.angelino/stack
```