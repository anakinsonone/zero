package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Add 2 numbers.",
	Long:    "Carry out addition operation on 2 numbers.",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err, res := Add(args[0], args[1], shouldRoundUp)

		if err != nil {
			return err
		}

		fmt.Printf("Addition of %s and %s = %s.\n\n", args[0], args[1], res)
		return nil
	},
}

func init() {
	addCmd.Flags().
		BoolVarP(&shouldRoundUp, "round", "r", false, "Round up result upto 2 decimal places.")
	rootCmd.AddCommand(addCmd)
}
