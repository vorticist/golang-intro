package schedule

import "fmt"

func main() {
	fmt.Println("main from package schedule")
}

// go build cmd api main
// go build cmd schedule main

// schedule lidia con cosas de schedule (tabla schedule_items)

// api main
// 	inicia server de API, sin referencia de nivel de aplicación (no request, no bd, no handlers)
// 	importar proyecto de handlers en main, y referenciar handlers
// repository conectarse a la BD
// service sirve para el Request
// controller están los handlers
