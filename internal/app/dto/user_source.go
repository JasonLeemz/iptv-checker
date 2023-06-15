package dto

type UserSourceDTO struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nick_name"`

	SourceName string `json:"source_name"`
	SourceURL  string `json:"source_url"`
	SourceData string `json:"source_data"`

	Ctime string `json:"ctime"`
	Utime string `json:"utime"`
}
