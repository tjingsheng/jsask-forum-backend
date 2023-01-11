package dataaccess

// import (
// 	"github.com/tjingsheng/jsask-forum-backend/internal/database"
// 	"github.com/tjingsheng/jsask-forum-backend/internal/models"
// )

// func ListComments(db *database.Database) (models.Comment, error) {
// 	comments := models.Comment{

// 		Post: models.Post{
// 			Post_ID:       5,
// 			User_ID:       1,
// 			Post_Datetime: "12 December 16:40",
// 			Post_Title:    "This is my fifth forum post",
// 			Post_Content:  "This is life 5 by one",
// 			Parent_Post:   -1,

// 			Tags: []string{"tag 1", "tag 3"},

// 			Username: "Bobby Lee One",

// 			Comment_Count: 23,

// 			Likes:               29,
// 			Is_Like_Selected:    true,
// 			Is_Dislike_Selected: false,
// 		},

// 		Comments: []models.Post{
// 			{
// 				Post_ID:       6,
// 				User_ID:       4,
// 				Post_Datetime: "12 December 17:40",
// 				Post_Title:    "NONE",
// 				Post_Content:  "1 Comment on this 5 by one",
// 				Parent_Post:   5,

// 				Tags: []string{},

// 				Username: "Harry Potter One",

// 				Comment_Count: 5,

// 				Likes:               10,
// 				Is_Like_Selected:    false,
// 				Is_Dislike_Selected: true,
// 			},
// 			{
// 				Post_ID:       7,
// 				User_ID:       6,
// 				Post_Datetime: "12 December 19:40",
// 				Post_Title:    "NONE",
// 				Post_Content:  "2 Comment on this 5 by Two",
// 				Parent_Post:   -1,

// 				Tags: []string{},

// 				Username: "Harry Potter Two",

// 				Comment_Count: 5,

// 				Likes:               999,
// 				Is_Like_Selected:    true,
// 				Is_Dislike_Selected: false,
// 			},
// 			{
// 				Post_ID:       8,
// 				User_ID:       4,
// 				Post_Datetime: "12 December 19:40",
// 				Post_Title:    "NONE",
// 				Post_Content:  "3 Comment on this 5 by one",
// 				Parent_Post:   5,

// 				Tags: []string{},

// 				Username: "Harry Potter One",

// 				Comment_Count: 5,

// 				Likes:               8,
// 				Is_Like_Selected:    true,
// 				Is_Dislike_Selected: false,
// 			},
// 			{
// 				Post_ID:       9,
// 				User_ID:       6,
// 				Post_Datetime: "12 December 20:40",
// 				Post_Title:    "NONE",
// 				Post_Content:  "4 Comment on this 5 by Two",
// 				Parent_Post:   -1,

// 				Tags: []string{},

// 				Username: "Harry Potter Two",

// 				Comment_Count: 5,

// 				Likes:               6,
// 				Is_Like_Selected:    false,
// 				Is_Dislike_Selected: false,
// 			},
// 		},
// 	}
// 	return comments, nil
// }
