package migrations

import (
	"github.com/sajeelwaien/gopro/database"
	"github.com/sajeelwaien/gopro/models"
)

func Migrate() {
	database.DBCon.AutoMigrate(models.Agent{})
}
