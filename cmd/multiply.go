package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	shouldRoundUp bool
	multiplyCmd   = &cobra.Command{
		Use:     "multiply",
		Aliases: []string{"mul", "multiple", "multi"},
		Short:   "Multiply 2 numbers.",
		Long:    "Carry out multiplication operation on two numbers.",
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(
				"Multiplication of %s and %s = %s.\n\n",
				args[0],
				args[1],
				Multiply(args[0], args[1], shouldRoundUp),
			)
		},
	}
)

func init() {
	multiplyCmd.Flags().
		BoolVarP(&shouldRoundUp, "round", "r", false, "Round result upto 2 decimal places.")
	rootCmd.AddCommand(multiplyCmd)
}
