package main

import (
	"fmt"
	
	"github.com/malvira/go-frontend"
)

var app = frontend.App{SourcePath: "./app"}//, DeployLocation: "somewhere/else.js"}

func main() {
	fmt.Println("test go-frontend")

	err := app.BuildAndServe(); if err != nil { fmt.Println("gopherjs build failed", err) }
}
