package main

import (
	"fmt"

	"github.com/ericsuh/template-repo-go/pkg/template"
)

func main() {
	fmt.Printf("worker: %s\n", template.Name())
}
