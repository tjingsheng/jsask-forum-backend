package dataaccess

import (
	"time"

	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"gorm.io/gorm"
)

func ListCurrUser(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	result := db.Where(models.User{Username: username}).Attrs(models.User{UserDatetime: time.Now().Format(constants.Go2PostgresqlDatetime)}).FirstOrCreate(&user)
	return user, result.Error
}
