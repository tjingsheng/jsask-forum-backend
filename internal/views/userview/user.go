package userview

import "github.com/tjingsheng/jsask-forum-backend/internal/models"

type ListView struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Username      string `json:"username"`
	User_Datetime string `json:"userDatetime"`
}

func ListFrom(user models.User) ListView {
	return ListView{
		ID:            user.ID,
		Username:      user.Username,
		User_Datetime: user.User_Datetime,
	}
}
