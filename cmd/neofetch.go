package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var neoFetchCmd = &cobra.Command{
	Use:   "neofetch",
	Short: "Show important system information",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		lipgloss.SetColorProfile(termenv.TrueColor)

		// baseStyle := lipgloss.NewStyle().
		// 	MarginTop(1).
		// 	MarginRight(4).
		// 	MarginBottom(1).
		// 	MarginLeft(4)

		var (
			b      strings.Builder
			lines  = strings.Split(art, "\n")
			colors = []string{"#CC3D00", "#CC3D00"}
			step   = len(lines) / len(colors)
		)

		for i, l := range lines {
			n := clamp(0, len(colors)-1, i/step)
			b.WriteString(colorize(colors[n], l))
			b.WriteRune('\n')
		}

		t := table.New().
			Border(lipgloss.HiddenBorder()).BorderStyle(lipgloss.NewStyle().Width(3))
		//TODO: show this specs
		// info := &specs.Specs{}
		// infoChan := make(chan specs.Specs, 1)
		// var wg sync.WaitGroup
		// wg.Add(1)
		// go getSpecs(info, infoChan, &wg)
		// wg.Wait()
		// newInfo := <-infoChan

		t.Row(b.String())

		fmt.Print(t)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(neoFetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func colorize(c, s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(c)).Render(s)
}

func clamp(v, low, high int) int {
	if high < low {
		low, high = high, low
	}
	return min(high, max(low, v))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
