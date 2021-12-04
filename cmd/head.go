package cmd

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func head() {
	myhead := figure.NewFigure("Shred", "isometric1", true).String()
	myhead = fmt.Sprintf("%s\n%s", myhead, "\tby Anton Feldmann")
	fmt.Printf(myhead + "\n\n\n\n")
}
