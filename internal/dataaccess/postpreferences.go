package dataaccess

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"gorm.io/gorm"
)

func UpdatePostPreference(db *gorm.DB, NewPreference models.PostPreference) error {
	result := db.Model(&NewPreference).Where(models.PostPreference{PostID: NewPreference.PostID, UserID: NewPreference.UserID}).UpdateColumn("preference", NewPreference.Preference)
	return result.Error
}
