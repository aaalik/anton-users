package migrate

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

var MigrateUpCommand = &cobra.Command{
	Use: "migrate-up",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := setupMigrateDB()
		if err != nil {
			log.Fatalf("[Migrate Up] Failed to setup migration: %v", err)
			return
		}

		err = m.Up()
		if err != nil {
			if err != migrate.ErrNoChange {
				log.Fatalf("[Migrate Up] Failed to migrate database: %v", err)
			}
			log.Printf("[Migrate Up] No more migrations to migrate")
			return
		}

		log.Println("[Migrate Up] Migrations successfully")
	},
}
