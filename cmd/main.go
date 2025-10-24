package main

import (
	"github.com/DimKa163/graph/app"
	"github.com/caarlos0/env"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var cfg app.Config

	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	server := app.NewServer(&cfg)
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
