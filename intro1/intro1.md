# Golang

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson (2009). It was originally created for use in cloud computing related tasks inside Google. It was designed with simplicity as one of the goals for the language and thus it is relatevily easy to get started with Go.

## References
- https://go.dev/tour/welcome/1
- https://www.calhoun.io/
- https://www.youtube.com/watch?v=etSN4X_fCnM&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM
- https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go
- https://www.youtube.com/c/JustForFunc

## Setup
- [Install instructions](https://go.dev/doc/install)
- [Goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/docs/languages/go)
- [Hello World](https://gobyexample.com/hello-world)
- [Tour of Go](https://go.dev/tour/welcome/1)
  
## Structs
Structs are the base of data management in Go as they are practical representations of the data that your program will process

```go
type Person struct {
    Name     string    // Upper-case at the start means this field will be exported     
    birthday time.Time // Lower-case means this field will not be exported
}
```

Go structs can also be annotated with what is know as a tag, in order to parse your structs into a desired output format or to make them work with databases

```go
type Car struct {
    VIN string `json:"vin_number" pg:"vin"` // {vin_number: ""}
}
```

- More: https://www.digitalocean.com/community/tutorials/defining-structs-in-go
- Suggested activity: Create a command line command that reads data from the user and stores it in a struc in memory, print back the result as a [JSON string](https://gobyexample.com/json).
  
## Functions as first class citizen
Functions are first class citizens in Go, that means that in our code we can treat functions as if the were any other data type
```go
func main() {
    m := "message"
    fn := func(p string) {
        fmt.Printf("input: %v\n", p)
    }
    fn(m)

    func(m) {
        fmt.Printf("input: %v", m)
    }(m)
}
```
```go
type MyFunc func(string) string
var f MyFunc

func main() {
    f = func(input string) string {
        return "output
    }
}
```
- More: https://golangbot.com/first-class-functions/
- Functional Programming: https://www.youtube.com/watch?v=KHojnWHemO0
## Receiver functions
Along with Interfaces, Receivers are the way to composition in Go. They allow to attach functions to types effectively giving structs behavior (methods).
```go
type Calculator struct {
    a int
    b int
}

func (c *Calculator) Sum() int {
    return c.a + c.b
}

func main() {
    calc := Calculator{a: 2, b: 5}
    fmt.Printf("result: %v\n", calc.Sum())
}
```
- More: https://golangbot.com/methods/
- Suggested activity
  - Create a program that takes input from the user
  - If the user entered plain text, bas64 encode the input and print back the result
  - If the user entered a base64 string, decode it and print the result back
  - Use a struct entity to process the users input
  - https://www.geeksforgeeks.org/base64-decodestring-function-in-golang-with-examples/ 
## Interfaces
Interfaces in Go allow composition, they are very flexible as they are implicit, meaning that a type does not need to declare itself as the type of the interface, but as long as it satisfies the the signature list in the interface, Go will consider it as that type.
```go
type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
``` 
- More: https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go

