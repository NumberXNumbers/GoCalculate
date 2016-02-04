package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//GoCalulateCmd is CoCaluate's root command
var GoCalulateCmd = &cobra.Command{
	Use:   "GoCalulate",
	Short: "GoCaulate is a command line tool for numerical analysis",
	Long: `Fast and Simple Numerical Analysis command line tool
            written in Go.
            Complete documentation is available at <website to come>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go Calculate World!")
	},
}
