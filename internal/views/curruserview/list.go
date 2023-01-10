package curruserview

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/curruserviewmodel"
)

type ListView struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Username     string `json:"username"`
	UserDatetime string `json:"userDatetime"`
}

func ListFrom(user curruserviewmodel.ListView) ListView {
	return ListView{
		ID:           user.ID,
		Username:     user.Username,
		UserDatetime: user.UserDatetime,
	}
}
