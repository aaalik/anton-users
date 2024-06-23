package main

import (
	"log"

	"github.com/aaalik/anton-users/cli/migrate"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{Use: "users"}

func init() {
	// Add more command here
	rootCommand.AddCommand(migrate.MigrateDownCommand)
	rootCommand.AddCommand(migrate.MigrateUpCommand)
	rootCommand.AddCommand(migrate.MirgateForceCommand)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		log.Panic(err)
	}
}
