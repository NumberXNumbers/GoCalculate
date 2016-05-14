package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//GoCalculateCmd is GoCalcuate's root command
var GoCalculateCmd = &cobra.Command{
	Use:   "GoCalulate",
	Short: "GoCaulate is a command line tool for numerical analysis",
	Long: `Fast and Simple Numerical Analysis command line tool
            written in Go.
            Complete documentation is available at <website to come>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go Calculate the World!")
	},
}

//Execute adds all child commands to the GoCalculate root command sets flags appropriately.
func Execute() {
	AddCommands()
	if err := GoCalculateCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// AddCommands adds child commands to the root command GoCalulateCmd.
func AddCommands() {
	GoCalculateCmd.AddCommand(calcCmd)
}
