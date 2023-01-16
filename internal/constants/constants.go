package constants

const (
	//Handler constants
	SuccessfulGetMessage    = "successfully GET data with %s"
	SuccessfulPutMessage    = "successfully PUT data with %s"
	SuccessfulPostMessage   = "successfully POST data with %s"
	SuccessfulDeleteMessage = "successfully DELETE data with %s"
	ErrRetrieveDatabase     = "failed to retrieve database in %s"
	ErrRetrieveData         = "failed to retrieve data in %s"
	ErrEncodeView           = "failed to retrieve data in %s"

	//Post Preference Enums
	IsLiked    = 1
	IsDisliked = 2

	//Datetime
	Go2PostgresqlDatetime = "2006-01-02 15:04:05-07"
)
