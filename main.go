package main

import (
	"fmt"

	"github.com/nicolas-calvario/ApiRest-Crud-Postgres/app"
)

func main() {
	fmt.Println("Crud api con postgres y gin")
	var a app.App
	a.CreateConnection()
	a.Routes()
}
