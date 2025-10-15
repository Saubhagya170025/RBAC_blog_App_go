package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// RunMigrations executes all migration files in the migrations directory
func RunMigrations(db *sql.DB, migrationsDir string) error {
	// Create migrations tracking table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations_applied (
			id SERIAL PRIMARY KEY,
			migration_name VARCHAR(255) NOT NULL UNIQUE,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Read migration files
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Sort files to ensure they run in order
	var migrationFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}
	sort.Strings(migrationFiles)

	// Execute each migration
	for _, migrationFile := range migrationFiles {
		applied, err := isMigrationApplied(db, migrationFile)
		if err != nil {
			return err
		}

		if applied {
			fmt.Printf("Migration %s already applied, skipping...\n", migrationFile)
			continue
		}

		// Read migration file
		filePath := filepath.Join(migrationsDir, migrationFile)
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", migrationFile, err)
		}

		// Execute migration
		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", migrationFile, err)
		}

		// Mark migration as applied
		_, err = db.Exec(
			"INSERT INTO migrations_applied (migration_name) VALUES ($1)",
			migrationFile,
		)
		if err != nil {
			return fmt.Errorf("failed to mark migration %s as applied: %w", migrationFile, err)
		}

		fmt.Printf("Migration %s executed successfully\n", migrationFile)
	}

	return nil
}

func isMigrationApplied(db *sql.DB, migrationName string) (bool, error) {
	var exists bool
	err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM migrations_applied WHERE migration_name = $1)",
		migrationName,
	).Scan(&exists)
	return exists, err
}