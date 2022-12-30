package tagsview

import "github.com/tjingsheng/jsask-forum-backend/internal/models"

type ListView struct {
	Tag_Name string `json:"tagName"`
}

func ListFrom(tag models.Tag) ListView {
	return ListView{
		Tag_Name: tag.Tag_Name,
	}
}
