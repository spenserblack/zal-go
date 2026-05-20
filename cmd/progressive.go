package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spenserblack/zal-go/corrupter"
	"github.com/spf13/cobra"
)

var (
	progressiveStartMin int
	progressiveStartMax int
	progressiveEndMin int
	progressiveEndMax int
)

var progressiveCmd = &cobra.Command{
	Use: "progressive <TEXT>",
	Short: "Progressively change the amount of corruption",
	Long: heredoc.Doc(`
		Progressively changes the level of corruption from the beginning to the end of
		text.
	`),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		stdout := cmd.OutOrStdout()
		stderr := cmd.ErrOrStderr()
		runes := []rune(args[0])

		err := assertMinMax(progressiveStartMin, progressiveStartMax)
		if err != nil {
			return err
		}
		err = assertMinMax(progressiveEndMin, progressiveEndMax)
		if err != nil {
			return err
		}
		corrupter := corrupter.New(stdout)
		lenFloat := float32(len(runes) - 1)
		for i, r := range runes {
			corrupter.Min = progressiveBlend(progressiveStartMin, progressiveEndMin, i, lenFloat)
			corrupter.Max = progressiveBlend(progressiveStartMax, progressiveEndMax, i, lenFloat)
			_, err := corrupter.WriteRune(r)
			if err != nil {
				fmt.Fprintln(stderr, err)
				return nil
			}
		}

		return nil
	},
}

// progressiveBlend helps get the current value between a start and end value based on
// the level of progress within the iteration. endPoint is the length
func progressiveBlend(start, end int, iteration int, endPoint float32) int {
	// NOTE Small optimization to skip calculations.
	if iteration == 0 {
		return start
	}
	diff := float32(end - start)
	progress := float32(iteration) / endPoint
	blend := float32(start) + (diff * progress)
	return progressiveRound(blend)
}

// progressiveRound rounds a float32, rounding to the nearest whole number.
func progressiveRound(f float32) int {
	if f < 0 {
		return int(f - 0.5)
	} else if f > 0 {
		return int(f + 0.5)
	} else {
		return 0
	}
}

func init() {
	progressiveCmd.PersistentFlags().IntVar(&progressiveStartMin, "start-min", 0, "The minimum corruption at the start of the text.")
	progressiveCmd.PersistentFlags().IntVar(&progressiveStartMax, "start-max", 1, "The maximum corruption at the start of the text.")
	progressiveCmd.PersistentFlags().IntVar(&progressiveEndMin, "end-min", 10, "The minimum corruption at the end of the text.")
	progressiveCmd.PersistentFlags().IntVar(&progressiveEndMax, "end-max", 20, "The maximum corruption at the end of the text.")
	rootCmd.AddCommand(progressiveCmd)
}
