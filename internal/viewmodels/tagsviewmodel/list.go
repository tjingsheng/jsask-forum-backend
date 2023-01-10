package tagsviewmodel

import "github.com/tjingsheng/jsask-forum-backend/internal/models"

type ListView struct {
	TagName string        `json:"tagName"`
	Posts   []models.Post `json:"posts" gorm:"many2many:posts_tags;joinForeignKey:id;"`
}

func ListFrom(tag models.Tag) ListView {
	return ListView{
		TagName: tag.TagName,
	}
}
