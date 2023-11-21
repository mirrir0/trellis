package main

import (
	"log"
	"os"

	arg "github.com/alexflint/go-arg"

	"github.com/31333337/bmrng/go/0kn/pkg/utils"
	"github.com/31333337/bmrng/go/trellis/client"
	"github.com/31333337/bmrng/go/trellis/config"
	"github.com/31333337/bmrng/go/trellis/errors"
	"github.com/31333337/bmrng/go/trellis/network"
	"github.com/31333337/bmrng/go/trellis/server"
)

// from cmd/server/server.go
func LaunchServer(args ArgsServer) {
	serversFile := args.ServerFile
	groupsFile := args.GroupFile
	addr := args.Addr
	errors.Addr = addr

	logger := utils.InitLogging()
	sugar := logger.Sugar()
	defer sugar.Sync()
	sugar.Infow(
		"Launching server",
		"address %s", addr,
	)

	servers, err := config.UnmarshalServersFromFile(serversFile)
	if err != nil {
		sugar.Fatalf("Could not read servers file %s", serversFile)
	}

	groups, err := config.UnmarshalGroupsFromFile(groupsFile)
	if err != nil {
		sugar.Fatalf("Could not read group file %s", groupsFile)
	}

	// will start in blocked state
	h := server.NewHandler()
	server := server.NewServer(&config.Servers{Servers: servers}, &config.Groups{Groups: groups}, h, addr)

	server.TcpConnections.LaunchAccepts()
	network.RunServer(h, server, servers, addr)
	config.Flush()
}

// from cmd/client/client.go
func LaunchClient(args ArgsClient) {
	serversFile := args.ServerFile
	groupsFile := args.GroupFile
	clientsFile := args.ClientFile
	addr := args.Addr
	errors.Addr = addr

	logger := utils.InitLogging()
	sugar := logger.Sugar()
	defer sugar.Sync()
	sugar.Infow(
		"Launching client",
		"address %s", addr,
	)

	servers, err := config.UnmarshalServersFromFile(serversFile)
	if err != nil {
		log.Fatalf("Could not read servers file %s", serversFile)
	}

	groups, err := config.UnmarshalGroupsFromFile(groupsFile)
	if err != nil {
		log.Fatalf("Could not read group file %s", groupsFile)
	}

	clients, err := config.UnmarshalServersFromFile(clientsFile)
	if err != nil {
		log.Fatalf("Could not read clients file %s", clientsFile)
	}

	clientRunner := client.NewClientRunner(servers, groups)

	// This fails @ network/rpc_call.go callee.HandleSignedMessageStream
	/*
		err = clientRunner.Connect()
		if err != nil {
			log.Fatalf("Could not make clients %v", err)
		}
	*/

	network.RunServer(nil, clientRunner, clients, addr)
}

func main() {
	var args Args
	argParser := arg.MustParse(&args)

	utils.SetDebugLogEnabled(args.Debug)
	utils.SetDebugLogCallerEnabled(args.DebugCaller)

	logger := utils.InitLogging()

	////////////////////////////////////////////////////////////////////////
	// change current working directory for trellis hardcoded paths TODO
	////////////////////////////////////////////////////////////////////////
	// TODO: handle this cleanly with a config file(?)
	trellisCoordinatorDirectory := "../../../trellis/cmd/coordinator/"
	if err := os.Chdir(trellisCoordinatorDirectory); err != nil {
		sugar := logger.Sugar()
		sugar.Fatalw(
			"Failed to change the working directory",
			"error: %v", err,
		)
		sugar.Sync()
		return
	}
	switch {
	case args.Coordinator != nil:
		LaunchCoordinator(*args.Coordinator, argParser, logger)

	case args.Client != nil:
		LaunchClient(*args.Client)

	case args.Server != nil:
		LaunchServer(*args.Server)

	default:
		argParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
}
