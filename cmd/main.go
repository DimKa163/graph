package main

import "github.com/DimKa163/graph/app"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	server := app.NewServer(&app.Config{
		Addr:     ":8082",
		Database: "postgres://postgres:NataZf0192274@localhost:5432/plan_date?sslmode=disable",
	})
	if err := server.AddServices(); err != nil {
		panic(err)
	}
	if err := server.AddLogging(); err != nil {
		panic(err)
	}
	server.Map()
	if err := server.Run(); err != nil {
		panic(err)
	}
}
