# Golang
## Channels
Channels are a device for communication between `goroutines`.
```go
var rc chan int
``` 
```go
rc := make(chan int)
```
Channels allow to pass data between `goroutines` by manipulating a reference to the channel from each of the `goroutines`. Channels will block when there is no value to read from the channel. You can send data to a channel when there is no data already in the channel, if a value has not been read from the channel when trying to send a new value, the execution will block until the value is read and the channel is free to receive more values.
```go
func process(c chan bool) {
    fmt.Printf("received: %v", <-c)
}

func main() {
    rc := make(chan bool)
    go process(rc)
    rc <- true
    fmt.Print("done...")
}
```
To close a channel we use the `close` function. When a channel is closed, all  of the `goroutines` that are trying to read from the channel will be notified to handle the close accordingly.
```go
func main() {
    c := make(chan int)
    defer close(c)
}

```
You can use the `for` loop to read values from a channel. If there are no values in the channel, the loop will block until there is a new value or the channel is closed.
```go
func main() {
    c := make(chan int)
    defer close(c)

    go func(c chan int) {
        for v := range c {
            fmt.Println(v)
        }
    }(c)

    for i := 0; i < 5; i++ {
        c <- i
    }
}
 
```
### Buffered Channels
Buffered channels are channels that can take a certain amount of values before blocking execution. Similarly, excution will not block when trying to read from a buffered channel unless there are no more values in the buffer.
```go
func main() {
    ch := make(chan int, 5)
    defer close(c)

    go func(c chan int) {
        for v := range c {
            fmt.Println(v)
        }
    }(c)

    for i := 0; i < 10; i++ {
        ch <- i
    }
}
```
### for select
The `select` statement is used to choose from several channel read/send operations. `select` will block execution while none of the operations can be resolved. It is commonly used along with the for loop as it allow continous monitoring of multiple channels.
```go
func process1(ch chan string) {
    for {  
        time.Sleep(3 * time.Second)
        ch <- "hello1"
    }
}
func process2(ch chan string) {
    for {  
        time.Sleep(2 * time.Second)
        ch <- "hello2"
    }

}
func main() {  
    output1 := make(chan string)
    output2 := make(chan string)
    go process1(output1)
    go process2(output2)
loop:
    for {
    select {
        case s1, ok := <-output1:
            if !ok {
                break loop
            }
            fmt.Println(s1)
        case s2 := <-output2:
            fmt.Println(s2)
        }
    }
}

``` 
- More
  - https://www.geeksforgeeks.org/channel-in-golang/
  - https://golangbot.com/select/
  - https://cyolo.io/blog/how-we-enabled-dynamic-channel-selection-at-scale-in-go/
- Suggested acitvity: Create a small Go application that reads rows from a database and parses them into Go structs. Parse each individual row on its own `goroutine` but synchronize the result using channels so that at the end we get a single list with all the parsed items.