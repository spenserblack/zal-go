package cmd

import (
	"fmt"
	"io"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spenserblack/zal-go/corrupter"
	"github.com/spf13/cobra"
)

var (
	// min is the minimum corruption.
	min int
	// max is the maximum corruption.
	max int
)

var rootCmd = &cobra.Command{
	Use:   "zalgo [flags] [string]",
	Short: "Corrupt text",
	Long: heredoc.Doc(`
		Corrupts text.

		If no text is passed, it is read from stdin.
	`),
	Example: "zalgo \"Hello, world!\"",
	Args:    cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		stdout := cmd.OutOrStdout()
		stderr := cmd.ErrOrStderr()
		stdin := cmd.InOrStdin()

		err := assertMinMax(min, max)
		if err != nil {
			return err
		}

		var bytes []byte
		if len(args) > 0 {
			bytes = []byte(args[0])
		} else {
			b, err := io.ReadAll(stdin)
			if err != nil {
				fmt.Fprintln(stderr, err)
				return nil
			}
			bytes = b
		}

		corrupter := corrupter.New(stdout)
		corrupter.Min = min
		corrupter.Max = max

		_, err = corrupter.Write(bytes)
		if err != nil {
			fmt.Fprintln(stderr, "Couldn't write: ", err)
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().IntVar(&min, "min", 1, "Set the minimum level of corruption")
	rootCmd.Flags().IntVar(&max, "max", 5, "Set the maximum level of corruption")
}
