package tagsviewmodel

import "github.com/tjingsheng/jsask-forum-backend/internal/models"

type ListView struct {
	TagName string `json:"tagName"`
}

func ListFrom(tag models.Tag) ListView {
	return ListView{
		TagName: tag.TagName,
	}
}
