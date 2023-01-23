package saltviewmodel

import "github.com/tjingsheng/jsask-forum-backend/internal/models"

type ListView struct {
	Username string `json:"username"`
	Salt     string `json:"salt"`
}

func ListFrom(user models.User) ListView {
	return ListView{
		Username: user.Username,
		Salt:     user.Salt,
	}
}
