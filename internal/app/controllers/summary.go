package controllers

import (
	"iptv-checker/core/context"
	"iptv-checker/internal/app/dto"
	"iptv-checker/internal/app/services"
	"strconv"
)

func GetUserSource(ctx ctx.Context) {
	userID, _ := strconv.Atoi(ctx.URLParam("user_id"))

	// Service 实例
	usService := services.NewUserSourceService()
	usDTO, err := usService.GetUserSource(int64(userID))
	var pd *userSource
	pd = convertPageData(usDTO)
	if err != nil {
		pd.Msg = err.Error()
	}

	ctx.JSON(pd)

}

type userSource struct {
	Code  int              `json:"code"`
	Msg   string           `json:"msg"`
	Count int              `json:"count"`
	Data  []userSourceData `json:"data"`
}

type userSourceData struct {
	SourceChannelName string `json:"source_channel_name"`
	SourceChannelUrl  string `json:"source_channel_url"`
	Ctime             string `json:"ctime"`
	Utime             string `json:"utime"`

	AuthorityId int    `json:"authorityId"`
	ParentId    int    `json:"parentId"`
	IsChannel   int    `json:"isChannel"`
	MenuIcon    string `json:"menuIcon"`
}

func convertPageData(us []*dto.UserSourceDTO) *userSource {
	pd := &userSource{
		Code:  0,
		Count: len(us),
		Data:  nil,
	}

	lineNum := 0
	authorityId := 0
	for i, v := range us {
		authorityId = i + 1 + lineNum
		pd.Data = append(pd.Data, userSourceData{
			SourceChannelName: v.SourceName,
			SourceChannelUrl:  v.SourceURL,
			Ctime:             v.Ctime,
			Utime:             v.Utime,

			AuthorityId: authorityId,
			ParentId:    -1,
			IsChannel:   0,
			MenuIcon:    "layui-icon-set",
		})
		// todo
		for _, channel := range v.SourceData {
			lineNum += 1

			pd.Data = append(pd.Data, userSourceData{
				SourceChannelName: channel.ChannelName,
				SourceChannelUrl:  channel.ChannelUrl,

				Ctime: channel.Ctime,
				Utime: channel.Utime,

				AuthorityId: authorityId + 1,
				ParentId:    authorityId,
				IsChannel:   1,
			})
		}
	}

	return pd
}
