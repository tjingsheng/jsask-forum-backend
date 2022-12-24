package handlers

const (
	ListUsers = "users.HandleListUsers"
	ListPosts = "users.HandleListPosts"
	ListTags  = "users.HandleListTags"

	SuccessfulListMessage = "successfully listed users"
	ErrRetrieveDatabase   = "failed to retrieve database in %s"
	ErrRetrieveUsers      = "failed to retrieve users in %s"
	ErrEncodeView         = "failed to retrieve users in %s"
)
