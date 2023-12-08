package migration

import (
	"github.com.br/rodrygo262/hackaton_golang/db"
	"github.com.br/rodrygo262/hackaton_golang/models"
)

func AutoMigration() {
	db := db.Connect()
	defer db.Close()

	db.Debug().AutoMigrate(&models.Provisioning{})
}
