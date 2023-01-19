package dataaccess

import (
	"time"

	"github.com/tjingsheng/jsask-forum-backend/internal/constants"
	"github.com/tjingsheng/jsask-forum-backend/internal/models"
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
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

func CleanPostPreferences(db *gorm.DB) {
	var postIds []int
	var jPostIds []int
	db.Model(&models.Post{}).Pluck("id", &postIds)
	db.Model(&models.PostPreference{}).Pluck("post_id", &jPostIds)
	db.Where("id IN ?", utils.SetDifference(jPostIds, postIds)).Delete(&models.PostPreference{})

	var userIds []int
	var jUserIds []int
	db.Model(&models.User{}).Pluck("id", &userIds)
	db.Model(&models.PostPreference{}).Pluck("user_id", &jUserIds)
	db.Where("id IN ?", utils.SetDifference(jUserIds, userIds)).Delete(&models.PostPreference{})
}
