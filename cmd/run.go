package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// this is the "nestri run" subcommand, takes no arguments for now
var runCmd = &cobra.Command{
	Use:   "run [options] [game]",
	Short: "Run a game using nestri",
	Args:  cobra.MaximumNArgs(1),
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
		// var gameConfig GameConfig
		// if err := viper.UnmarshalKey(fmt.Sprintf("games.%s", game), &gameConfig); err != nil {
		// 	return fmt.Errorf("error parsing game configuration: %w", err)
		// }

		// // Override config with command-line flags
		// cmd.Flags().Visit(func(f *pflag.Flag) {
		// 	switch f.Name {
		// 	case "directory":
		// 		gameConfig.Directory = viper.GetString(f.Name)
		// 	case "executable":
		// 		gameConfig.Executable = viper.GetString(f.Name)
		// 	case "gpu":
		// 		gameConfig.GPU = viper.GetInt(f.Name)
		// 	case "vendor":
		// 		gameConfig.Vendor = viper.GetString(f.Name)
		// 	case "resolution-height":
		// 		gameConfig.Resolution.Height = viper.GetInt(f.Name)
		// 	case "resolution-width":
		// 		gameConfig.Resolution.Width = viper.GetInt(f.Name)
		// 	}
		// })

		fmt.Printf("Running game: %s\n\n", game)

		// cli, err := client.NewClientWithOpts(client.FromEnv)
		// if err != nil {
		// 	return fmt.Errorf("error creating Docker client: %w", err)
		// }

		// ctx := context.Background()

		// var game string
		// if len(args) > 0 {
		// 	game = args[0]
		// 	viper.Set("game", game)
		// 	viper.WriteConfig()
		// } else {
		// 	game = viper.GetString("game")
		// 	if filepath.Ext(game) != ".exe" {
		// 		return fmt.Errorf("Make sure the game is a .exe")
		// 	}
		// 	if game == "" {
		// 		return fmt.Errorf("no game specified and no previous game selected")
		// 	}
		// }

		// fmt.Printf("Running game: %s\n\n", game)

		// cli, err := client.NewClientWithOpts()
		// if err != nil {
		// 	panic(err)
		// }

		// ctx := context.Background()
		// resp, err := cli.ContainerCreate(ctx, &container.Config{
		// 	Image: "hello-world",
		// }, nil, nil, nil, "hello-world")

		// if err != nil {
		// 	panic(err)
		// }

		// if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		// 	panic(err)
		// }

		// // Attach to the container to get logs
		// out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
		// if err != nil {
		// 	fmt.Printf("Error attaching to container logs: %s\n", err)
		// }
		// defer out.Close()

		// // Copy the logs to stdout and stderr
		// stdcopy.StdCopy(os.Stdout, os.Stderr, out)

		// // Wait for the container to finish
		// statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
		// select {
		// case err := <-errCh:
		// 	if err != nil {
		// 		fmt.Printf("Error waiting for container: %s\n", err)
		// 	}
		// case <-statusCh:
		// 	fmt.Println("Container finished")
		// }
		// // Clean up the container
		// if err := cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{}); err != nil {
		// 	fmt.Printf("Error removing container: %s\n", err)
		// }
		// if gpu > 0 {
		// 	fmt.Print("Using gpu %s\n", gpu)
		// }
		// if hdr {
		// 	fmt.Println("Enabling HDR mode")
		// }

		//get linux version
		// versionCmd := exec.Command("grep", "VERSION", "/etc/os-release")
		// versionOutput, err := versionCmd.CombinedOutput()
		// if err != nil {
		// 	return fmt.Errorf("error getting linux version:")
		// }
		// fmt.Printf("Linux version:\n%s\n", string(versionOutput))

		// //Step 1: change to games dir
		// fmt.Println("changing to game dir.") //this is a temp command for debug as well as leads to a hardcoded dir

		// HomeDir, err := os.UserHomeDir()
		// if err != nil {
		// 	return fmt.Errorf("error getting home directory %v\n", err)
		// }

		// err = os.Chdir(fmt.Sprintf("%s/game", HomeDir))
		// if err != nil {
		// 	return fmt.Errorf("error changing directory: %v\n", err)
		// }
		// //verify we are in game dir
		// dir, err := os.Getwd()
		// if err != nil {
		// 	return fmt.Errorf("error getting current directory: %v\n", err)
		// }
		// fmt.Printf("Current directory: %s\n\n", dir)

		// //list games dir
		// listDir := exec.Command("ls", "-la", ".")
		// listDirOutput, err := listDir.CombinedOutput()
		// if err != nil {
		// 	fmt.Errorf("error listing games: %v\n")
		// }
		// fmt.Printf("List of Games: \n%s\n", listDirOutput)

		// //step 2: Generate a Session ID
		// //generate id
		// SID := exec.Command("bash", "-c", "head /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | head -c 16")

		// //save output to variable
		// output, err := SID.Output()
		// if err != nil {
		// 	fmt.Errorf("Error generating Session ID: %v\n", err)
		// }
		// sessionID := strings.TrimSpace(string(output))
		// fmt.Printf("Your Session ID is: %s\n\n", sessionID)

		// //step 3: Launch netris server
		// fmt.Println("Installing Netris/Launching Netris Server")
		// checkRunning := exec.Command("sudo", "docker", "ps", "-q", "-f", "name=netris")
		// containerId, err := checkRunning.Output()
		// if err != nil {
		// 	return fmt.Errorf("error checking running Docker container: %v", err)
		// }

		// if len(containerId) == 0 {
		// 	checkExisting := exec.Command("sudo", "docker", "ps", "-aq", "-f", "name=netris")
		// 	containerId, err = checkExisting.Output()
		// 	if err != nil {
		// 		return fmt.Errorf("error checking for existing docker container: %v", err)
		// 	}

		// 	if len(containerId) == 0 {
		// 		installCmd := exec.Command(
		// 			"sudo", "docker", "run", "-d", "--gpus", "all", "--device=/dev/dri",
		// 			"--name", "netris", "-it", "--entrypoint", "/bin/bash",
		// 			"-e", fmt.Sprintf("SESSION_ID=%s", sessionID),
		// 			"-v", fmt.Sprintf("%s:/game", dir), "-p", "8080:8080/udp",
		// 			"--cap-add=SYS_NICE", "--cap-add=SYS_ADMIN", "ghcr.io/netrisdotme/netris/server:nightly",
		// 		)
		// 		installCmd.Stdout = os.Stdout
		// 		installCmd.Stderr = os.Stderr

		// 		if err := installCmd.Run(); err != nil {
		// 			return fmt.Errorf("error running docker command: %v", err)
		// 		}
		// 	} else {
		// 		startContainer := exec.Command("sudo", "docker", "start", "netris")
		// 		startContainer.Stdout = os.Stdout
		// 		startContainer.Stderr = os.Stderr

		// 		if err := startContainer.Run(); err != nil {
		// 			return fmt.Errorf("error starting existing Docker container: %v", err)
		// 		}
		// 	}
		// }

		// //main part of step 4:
		// //start netris server

		// fmt.Println("starting netris server\n\n")
		// checkFileCmd := exec.Command("sudo", "docker", "exec", "netris", "ls", "-la", "/tmp")
		// output, err = checkFileCmd.Output()
		// if err != nil {
		// 	return fmt.Errorf("error checking /tmp dir in docker container: %v\n", err)
		// }

		// if !strings.Contains(string(output), ".X11-unix") {
		// 	startupCmd := exec.Command("sudo", "docker", "exec", "netris", "/etc/startup.sh", ">", "/dev/null", "&")
		// 	startupCmd.Stdout = os.Stdout
		// 	startupCmd.Stderr = os.Stderr

		// 	if err := startupCmd.Run(); err != nil {
		// 		return fmt.Errorf("error running startup command: %v\n", err)
		// 	}

		// 	for {
		// 		time.Sleep(7 * time.Minute)
		// 		output, err := checkFileCmd.Output()
		// 		if err != nil {
		// 			return fmt.Errorf("error checking /tmp directory in container: %v\n", err)
		// 		}
		// 		if strings.Contains(string(output), ".X11-unix") {
		// 			break
		// 		}
		// 	}
		// }

		// gameCmd := fmt.Sprintf("netris-proton -pr %s", game)
		// execCmd := exec.Command("sudo", "docker", "exec", "netris", gameCmd)
		// execCmd.Stdout = os.Stdout
		// execCmd.Stderr = os.Stderr

		// if err := execCmd.Run(); err != nil {
		// 	return fmt.Errorf("error executing game command in docker container: %v\n", err)
		// }

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
