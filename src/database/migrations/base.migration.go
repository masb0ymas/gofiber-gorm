package migrations

import "os"

func BaseMigration() string {
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_TIMEZONE := os.Getenv("DB_TIMEZONE")

	schema := `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		SELECT * FROM pg_timezone_names;

		ALTER DATABASE ` + DB_DATABASE + ` SET timezone TO '` + DB_TIMEZONE + `';
	`

	return schema
}
