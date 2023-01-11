package currpostsview

import (
	"github.com/tjingsheng/jsask-forum-backend/internal/utils"
	"github.com/tjingsheng/jsask-forum-backend/internal/views/postsview"
)

type ListView struct {
	Posts    postsview.ListView   `json:"post"`
	Comments []postsview.ListView `json:"comments"`
}

func isPost(post postsview.ListView) bool {
	return post.ParentPost == 0
}

func isComment(post postsview.ListView) bool {
	return post.ParentPost != 0
}

func ListFrom(post []postsview.ListView) ListView {
	currPost := utils.FilterStruct(post, isPost)[0]
	comments := utils.FilterStruct(post, isComment)

	return ListView{
		Posts:    currPost,
		Comments: comments,
	}
}
