package main

import (
	_ "github.com/YReshetko/go-annotation/annotations/constructor"
	_ "github.com/YReshetko/go-annotation/annotations/mock"
	_ "github.com/YReshetko/go-annotation/annotations/validator"
	"github.com/YReshetko/go-annotation/pkg"
)

func main() {
	annotation.Process()
}
