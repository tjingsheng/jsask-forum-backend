package handlers

const (
	HandleList      = "HandleList"
	ListComments    = "Comments.HandleListComments"
	ListPosts       = "Posts.HandleListPosts"
	ListTags        = "Tags.HandleListTags"
	ListCurrentUser = "Tags.HandleListTags"

	SuccessfulListMessage = "successfully listed data with %s"
	ErrRetrieveDatabase   = "failed to retrieve database in %s"
	ErrRetrieveData       = "failed to retrieve data in %s"
	ErrEncodeView         = "failed to retrieve data in %s"
)
