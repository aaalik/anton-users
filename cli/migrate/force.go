package migrate

import (
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

var MirgateForceCommand = &cobra.Command{
	Use:  "migrate-force",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := setupMigrateDB()
		if err != nil {
			log.Fatalf("[Migrate Force] Failed to setup migration: %v", err)
			return
		}

		version, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("[Migrate Force] Invalid arg: %v", err)
			return
		}

		err = m.Force(version)
		if err != nil {
			if err != migrate.ErrNoChange {
				log.Fatalf("[Migrate Force] Failed to migrate database: %v", err)
			}
			log.Printf("[Migrate Force] No more migrations to migrate")
			return
		}

		log.Printf("[Migrate Force] Migrate to version %v has been successfully", version)
	},
}
