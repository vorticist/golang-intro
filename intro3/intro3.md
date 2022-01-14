# Golang
## Go Modules
Go Modules is the way Go has to handle dependency management. It provides explicit commands to manage dependencies, but in addition, as part of the `go build` and `go test` processes, dependencies will be solved automatically when running those commands.
### go mod
This command is the main point of interaction for dependency management. 

#### go mod init
A go module is initialized with the `go mod init` command.
```bash
go mod init github.com/vorticist/intro3
```
After running this command two file will be generated, `go.mod` and `go.sum`
- `go.mod` is the file where all dependencies will be listed, at first it will be almost empty but as soon as you start adding dependencies, they'll be added to this file, it is from here that `go mod` knows which dependencies to fetch when building or running `go mod download`
  ```go
  module github.com/vorticist/intro3

  go 1.17

  require (
      github.com/aws/aws-lambda-go v1.27.0
	  github.com/golang/protobuf v1.5.2 // indirect
  )
  ```
- `go.sum` this file contains a hash, for each dependency, that is used to validate the code that is downloaded by `go mod`.

#### go mod tidy
`go mod tidy` examines the code and cleans up (tidies) the `go.mod` file accordingly, It removes unused references and adds them when new ones are found in the code. It will also update the `go.sum` file. You can add dependencies in your code and then run `go mod tidy` to update the `go.mod` file
```go
import "github.com/aws/aws-lambda-go"

func main() {

}
``` 
```bash
go mod tidy
```
```go
module github.com/vorticist/intro3

go 1.17

require (
    github.com/aws/aws-lambda-go v1.27.0
)
```
#### go mod download
When used without arguments, this command downloads all the dependencies specified in the `go.mod` file. It creates a cache folder for each of the modules downloaded this way. 
```bash
go mod download
```
When passed as a parameter `go mod download` will only download the specified module.
```bash
go mod download github.com/aws/aws-lambda-go v1.27.0
```
### go install
Builds and installs the module indicated by the path
```bash
go install github.com/aws/aws-lambda-go@v1.27.0
```
### go get
Updates the `go.mod` references and then builds and installs the module indicated by the path
```bash
go get github.com/aws/aws-lambda-go
```
- More: https://go.dev/ref/mod
### Packages
In Go we have the main module, which corresponds to the root folder of the project, each of the subfolders is considered its own package and is accesible via the main module package name + the path to the folder, so for instance in the sample below, if our main module is named `github.com/vorticist/intro3` then our `internal` folder would be accesible by referencing `github.com/vorticist/intro3/internal` and the `api` package would be `github.com/vorticist/intro3/cmd/api`
```bash
├── cmd
│   ├── api
│   │   └── main.go
│   └── cron
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── bootstrap
│   │   ├── api.go
│   │   └── cron.go
│   ├── config
│   │   └── config.go
│   ├── controller
│   │   ├── controller.go
│   │   └── metrics.go
│   ├── infraestructure
│   │   ├── keycloak.go
│   │   └── postgres.go
│   ├── models
│   │   └── models.go
│   ├── repository
│   │   ├── coding.go
│   │   ├── delivery.go
│   │   ├── deployment.go
│   │   ├── metrics.go
│   │   └── repository.go
│   └── service
│       ├── coding.go
│       ├── delivery.go
│       ├── deployment.go
│       ├── metrics.go
│       └── service.go
├── Makefile

```
- More: https://www.golang-book.com/books/intro/11
- Suggested activity: Using `gorilla` create a small API application that receives a request with a string message and writes that message into a file. Try to organize your code into packages that make sense for the steps involved in this process. 
## Go routines
Consider them as lightweight threads managed by the Go runtime. Goroutines are one of the central tools provided by the language to work with concurrency.
```go
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```
### sync.WaitGroup
One of the ways that we have to synchronize concurrent executions of goroutines. As the name suggest a `WaitGroup` will block execution until all members of the group are done doing their work.
```go
func main() {
  var wg sync.WaitGroup

  process := func(item string) {
    fmt.Printf("processing %v", item)
  }
  items := string{"item1", "item2", "item3"}

  wg.Add(len(items))
  for _, item := range items {
    go func() {
      defer wg.Done()
      process(item)
    }()
  }

  wg.Wait()
}
```
- More: https://www.golangprograms.com/goroutines.html