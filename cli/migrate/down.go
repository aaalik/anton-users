package migrate

import (
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

var MigrateDownCommand = &cobra.Command{
	Use:  "migrate-down",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := setupMigrateDB()
		if err != nil {
			log.Fatalf("[Migrate Down] Failed to setup migration: %v", err)
			return
		}

		steps, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("[Migrate Down] Invalid arg: %v", err)
			return
		}

		// using loop for better error handling perstep, and adaptability
		// this won't make performance overhead because step < 10000
		for i := 0; i < steps; i++ {
			err = m.Steps(-1)
			if err != nil {
				if err != migrate.ErrNoChange {
					log.Fatalf("[Migrate Down] Failed to rollback database: %v", err)
				}
				log.Printf("[Migrate Down] No more migration to rollback")
				break
			}
			log.Println("[Migrate Down] Successfully rollback one step of migration")
		}
	},
}
