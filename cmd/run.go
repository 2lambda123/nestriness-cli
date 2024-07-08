package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// this is the "nestri run" subcommand, takes no arguments for now
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a game using nestri",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if runtime.GOOS != "linux" {
			//make sure os is linux
			fmt.Println("This command is only supported on Linux.")
			return nil
		}

		//The main job here is to:
		//1. look for an exe in a certain directory,
		//2. mount the directory inside the container
		//3. Run the nestri docker container
		//4. SSH into the container and set up everything
		//5. Run the game
		//6. Provide the URL to play or throw an error otherwise.

		// The last argument is the game to run
		game := args[len(args)-1]

		// Load game configuration
		var gameConfig GameConfig
		if err := viper.UnmarshalKey(fmt.Sprintf("games.%s", game), &gameConfig); err != nil {
			return fmt.Errorf("error parsing game configuration: %w", err)
		}

		flags := cmd.Flags()

		if flags.Changed("directory") || flags.Changed("d") {
			gameConfig.Directory, _ = flags.GetString("directory")
		}
		if flags.Changed("executable") || flags.Changed("x") {
			gameConfig.Executable, _ = flags.GetString("executable")
		}
		if flags.Changed("gpu") {
			gameConfig.GPU, _ = flags.GetInt("gpu")
		}
		if flags.Changed("vendor") || flags.Changed("v") {
			gameConfig.Vendor, _ = flags.GetString("vendor")
		}
		if flags.Changed("height") || flags.Changed("H") {
			gameConfig.Resolution.Height, _ = flags.GetInt("height")
		}
		if flags.Changed("width") || flags.Changed("W") {
			gameConfig.Resolution.Width, _ = flags.GetInt("width")
		}

		fmt.Println("Game config:", gameConfig)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("directory", "d", "", "Game directory")
	runCmd.Flags().StringP("executable", "x", "", "Game executable")
	runCmd.Flags().Int("gpu", 0, "GPU number")
	runCmd.Flags().StringP("vendor", "v", "", "GPU vendor")
	runCmd.Flags().IntP("height", "H", 1080, "Screen height")
	runCmd.Flags().IntP("width", "W", 1920, "Screen width")

	// viper.BindPFlag("directory", runCmd.Flags().Lookup("directory"))
	// viper.BindPFlag("executable", runCmd.Flags().Lookup("executable"))
	// viper.BindPFlag("gpu", runCmd.Flags().Lookup("gpu"))
	// viper.BindPFlag("vendor", runCmd.Flags().Lookup("vendor"))
	// viper.BindPFlag("resolution.height", runCmd.Flags().Lookup("height"))
	// viper.BindPFlag("resolution.width", runCmd.Flags().Lookup("width"))

	viper.BindPFlag("games.*.directory", runCmd.Flags().Lookup("directory"))
	viper.BindPFlag("games.*.executable", runCmd.Flags().Lookup("executable"))
	viper.BindPFlag("games.*.gpu", runCmd.Flags().Lookup("gpu"))
	viper.BindPFlag("games.*.vendor", runCmd.Flags().Lookup("vendor"))
	viper.BindPFlag("games.*.resolution.height", runCmd.Flags().Lookup("height"))
	viper.BindPFlag("games.*.resolution.width", runCmd.Flags().Lookup("width"))
}
