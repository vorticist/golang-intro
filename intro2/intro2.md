# Golang

## HTTP
`net/http` is a powerful, yet really simple to use, package in Go that allows developers to serve code over the net. More complex packages and frameworks are modeled after th `net/http` package.
```go
func main() {
    http.HanldelFunc("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "hello world!")
    })

    http.ListenAndServe(":8080", nil)
}
```
- More: https://codegangsta.gitbooks.io/building-web-apps-with-go/content/http_basics/index.html

### Third party packages
Although the `net/http` package is pretty complete, several third party libraries and packages have come up to improve the developer experience. Two of the most popular ones are `gorilla` and `echo`   

#### Gorilla
This one is more a toolkit rather than a library. Its router is accepts handlers compatible with the `net/http` package, so it's really easy to add `gorilla` to an existing app. Also has tools to handle using things like websockets, cookies, rpcs, etc.

```go
import "github.com/gorilla/mux"

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)
    http.Handle("/", r)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
    // Do someting
}
```
- More: [Gorilla](https://www.gorillatoolkit.org/)
- Suggested activities: Create a simple [API server application using gorilla](https://www.soberkoder.com/go-rest-api-gorilla-mux/) 

### Echo
A small web framework focused on high performance. It encompases the same funcitonality as the `net/http` package, but with its own types and interfaces.
```go
import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```
- More: https://echo.labstack.com/guide/
- Create a small [API server using echo](https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508)

## Using a database
Similar to other technologies if you want to connect to any particular database engine you're gonna need a driver package either proprietary or open source.
```go
import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

func main() {
    ...

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()
}
```
- Popular drivers; [Postgres](https://github.com/lib/pq), [mongodb](https://docs.mongodb.com/drivers/go/current/), [mysql](https://github.com/go-sql-driver/mysql)

A common pattern for interacting with databases is to use a repository
```go
type UsersRepository interface {
    New(user User)
    Update(user User)
    Delete(userId string)
    List() []User
}

type userRepository struct {
    db *sql.DB
}

func (ur *userRepository) New(user User) {
    ...
}
```
- More 
  - https://threedots.tech/post/repository-pattern-in-go/
  - https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
  - https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
- Suggested activity: Create a small API project that saves requests data into a database of your choice

## Building Apps
`go build` is the tool provided to compile our go applications and package them into a binary. But `go build` also accepts several flags to indicate the target platform of the binary and ouput destination. Meaning you can write your code in one OS, but have the binary target a different platform.
```shell
GOOS=darwin GOARCH=amd64 go build .
```
This command will build a project into a binary that can be run only on macs

- More:
  - https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures
  - [List of supported OS and architechtures](https://go.dev/doc/install/source#environment)
- Suggested activities: Build a small go app into a binary that canbe ran in a Raspberry Pi