package main

import (
	"log"
	"os"
	"path/filepath"

	arg "github.com/alexflint/go-arg"

	"go.uber.org/zap"

	"github.com/31333337/bmrng/go/0kn/pkg/utils"
)

func setWorkingDirectory() {
	// get working directory from env var, ensure set for children processes
	workingDir := os.Getenv("_0KN_WORKDIR")
	if workingDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic("Error getting user's home directory: " + err.Error())
		}
		defaultDir := filepath.Join(homeDir, ".0KN")

		log.Println("Env _0KN_WORKDIR not set, using default", defaultDir)
		workingDir = defaultDir
		os.Setenv("_0KN_WORKDIR", workingDir)
	}

	// create working directory if it does not exist
	subDirs := []string{
		"certificates",
	}
	for _, d := range subDirs {
		dir := workingDir + "/" + d
		if err := os.MkdirAll(dir, os.FileMode(0700)); err != nil {
			panic("Failed to create directory: " + dir + "; " + err.Error())
		}
	}

	// change to working directory
	if err := os.Chdir(workingDir); err != nil {
		panic("Failed to change the working directory: " + err.Error())
	}
}

func getWorkingDirectory() string {
	return os.Getenv("_0KN_WORKDIR")
}

func main() {
	var args Args
	argParser := arg.MustParse(&args)

	utils.SetDebugLogEnabled(args.Debug)
	utils.SetDebugLogCallerEnabled(args.DebugCaller)

	setWorkingDirectory()

	switch {
	case args.Coordinator != nil:
		LaunchCoordinator(*args.Coordinator, argParser, logger)

	case args.Client != nil:
		LaunchClient(*args.Client, logger)

	case args.Server != nil:
		LaunchServer(*args.Server, logger)

	default:
		argParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
}
