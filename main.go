package main

import (
	"Exe/Model"
	"Exe/cmd"
)

func main() {
	Model.InitDB()
	cmd.InitRouter()
}
