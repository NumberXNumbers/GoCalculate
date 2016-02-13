package main

import (
	"fmt"
	"os"

	"github.com/NumberXNumbers/GoCalculate/cmd"
)

func main() {
	if err := cmd.GoCalulateCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
