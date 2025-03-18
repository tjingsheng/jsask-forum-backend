package tagsview

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/viewmodels/tagsviewmodel"
)

type ListView struct {
	Value string `json:"value"`
	Label string `json:"name"`
}

func ListFrom(tag tagsviewmodel.ListView) ListView {
	return ListView{
		Value: tag.TagName,
		Label: tag.TagName,
	}
}
