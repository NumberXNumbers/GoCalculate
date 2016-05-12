package cmd

import (
	"fmt"
	"os"

	"github.com/NumberXNumbers/GoCalculate/calculators"
	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/spf13/cobra"
)

func init() {
	calcCmd.AddCommand(infixCmd)
	calcCmd.AddCommand(postfixCmd)
	calcCmd.AddCommand(prefixCmd)
}

func calcErr(method string, e error) {
	fmt.Printf("An Error Occured in method %s: %v", method, e)
	os.Exit(-1)
}

// This will hold code to calculate using infix, prefix or postfix notation
var calcCmd = &cobra.Command{
	Use:   "calculator",
	Short: "Simple calculator tool for simple calculations",
	Long: `Simple calculator tool for simple calculations.
    Requires sub-commands infix, postfix or prefix`,
}

var infixCmd = &cobra.Command{
	Use:   "infix",
	Short: "Simple Infix calculator",
	Long: `Simple Infix calculator.
    All inputs must be of the in the prefix notation with spaces between each value and/or operation
    Example '( 4 + 3 )' results in 7`,
	Run: func(cmd *cobra.Command, args []string) {
		value, err := calculators.InfixCalculator(args)
		if err != nil {
			calcErr("InfixCalculator", err)
		}
		if value.Type() == gcv.Complex {
			fmt.Println(value.Complex())
		} else {
			fmt.Println(value.Real())
		}
	},
}

var postfixCmd = &cobra.Command{
	Use:   "postfix",
	Short: "Simple Postfix calculator",
	Long: `Simple Postfix calculator.
    All inputs must be of the in the postfix notation with spaces between each value and/or operation
    Example '4 3 +' results in 7`,
	Run: func(cmd *cobra.Command, args []string) {
		value, err := calculators.ReversePolishCalculator(args)
		if err != nil {
			calcErr("ReversePolishCalculator", err)
		}
		if value.Type() == gcv.Complex {
			fmt.Println(value.Complex())
		} else {
			fmt.Println(value.Real())
		}
	},
}

var prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Simple Prefix calculator",
	Long: `Simple Prefix calculator.
    All inputs must be of the in the prefix notation with spaces between each value and/or operation
    Example '+ 4 3' results in 7`,
	Run: func(cmd *cobra.Command, args []string) {
		value, err := calculators.PolishCalculator(args)
		if err != nil {
			calcErr("PolishCalculator", err)
		}
		if value.Type() == gcv.Complex {
			fmt.Println(value.Complex())
		} else {
			fmt.Println(value.Real())
		}
	},
}
