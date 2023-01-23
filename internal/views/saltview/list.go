package saltview

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/saltviewmodel"
)

type ListView struct {
	Username string `json:"username"`
	Salt     string `json:"salt"`
}

func ListFrom(user saltviewmodel.ListView) ListView {
	return ListView{
		Username: user.Username,
		Salt:     user.Salt,
	}
}
