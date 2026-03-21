package main

import (
	"Project03-Go_Bootcamp/internal/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(di.Module).Run()
}
