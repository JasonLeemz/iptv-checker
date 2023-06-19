package dto

type UserSourceDTO struct {
	UserSourceID int64 `json:"user_source_id"`
	LiveSourceID int64 `json:"live_source_id"`

	UserID   int64  `json:"user_id"`
	NickName string `json:"nick_name"`

	SourceName string `json:"source_name"`
	SourceURL  string `json:"source_url"`
	SourceData string `json:"source_data"`

	Ctime string `json:"ctime"`
	Utime string `json:"utime"`
}
