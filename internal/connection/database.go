package connection

import (
	"fmt"
	"log"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/config"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/models"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(config config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=%s", 
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
		config.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database : ", err)
	}

	migration := models.Migration{}
	if err = db.AutoMigrate(migration.RegisterModels()...); err != nil {
		utils.Log.Error("Failed to migrate database : ", "Error", err)
	}

	return db
}