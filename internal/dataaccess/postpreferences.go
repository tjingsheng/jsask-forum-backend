package dataaccess

import (
	"time"

	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"gorm.io/gorm"
)

func UpdatePostPreference(db *gorm.DB, NewPreference models.PostPreference) error {
	result := db.Model(&NewPreference).
		Where(models.PostPreference{PostID: NewPreference.PostID, UserID: NewPreference.UserID}).
		Assign(models.PostPreference{
			Preference:         NewPreference.Preference,
			PreferenceDatetime: time.Now().Format(constants.Go2PostgresqlDatetime)}).
		FirstOrCreate(&NewPreference)
	return result.Error
}
