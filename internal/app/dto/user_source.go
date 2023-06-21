package dto

type UserSourceDTO struct {
	UserSourceID int64 `json:"user_source_id"`
	LiveSourceID int64 `json:"live_source_id"`

	UserID   int64  `json:"user_id"`
	NickName string `json:"nick_name"`

	SourceName string          `json:"source_name"`
	SourceURL  string          `json:"source_url"`
	SourceData []ChannelSource `json:"source_data" gorm:"-"`

	Ctime string `json:"ctime"`
	Utime string `json:"utime"`
}

type ChannelSource struct {
	ChannelID   int64  `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ChannelUrl  string `json:"channel_url"`
	Ctime       string `json:"ctime"`
	Utime       string `json:"utime"`
}
