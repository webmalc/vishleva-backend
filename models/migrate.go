package models

// Migrate migrates the DB
func Migrate(migrater AutoMigrater) {
	migrater.AutoMigrate(
		&User{},
		&Collection{},
		&Tag{},
		&Image{},
		&Tariff{},
	)
}
