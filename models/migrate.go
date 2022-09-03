package models

// Migrate migrates the DB.
func Migrate(migrater AutoMigrater) {
	migrater.AutoMigrate(
		&User{},
		&Collection{},
		&Tag{},
		&Image{},
		&Tariff{},
		&Review{},
		&Client{},
		&Order{},
	)
	migrater.Model(&Order{}).AddForeignKey(
		"client_id", "clients(id)", "RESTRICT", "CASCADE",
	)
	migrater.Model(&Review{}).AddForeignKey(
		"client_id", "clients(id)", "SET NULL", "CASCADE",
	)
	migrater.Model(&Review{}).AddForeignKey(
		"image_id", "images(id)", "SET NULL", "CASCADE",
	)
	migrater.Model(&Collection{}).AddForeignKey(
		"client_id", "clients(id)", "SET NULL", "CASCADE",
	)
}
