package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"sub"},
	Short:   "Subtract a number from another.",
	Long:    "Carry out subtraction operation on 2 integers.",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err, res := Subtract(args[0], args[1], shouldRoundUp)

		if err != nil {
			return err
		}

		fmt.Printf("Subtraction of %s from %s = %s.\n\n", args[0], args[1], res)
		return nil
	},
}

func init() {
	subractCmd.Flags().
		BoolVarP(&shouldRoundUp, "round", "r", false, "Round up result upto 2 decimal places.")
	rootCmd.AddCommand(subractCmd)
}
