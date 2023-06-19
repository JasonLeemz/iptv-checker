package logic

import (
	"iptv-checker/internal/app/services"
)

func CheckCanDelSource(userID int64, userSourceID []string) bool {
	// 如果查到的结果和userSourceID长度一致，返回true
	userSourceSrv := services.NewUserSourceService()
	cond := map[string]interface{}{
		"user_id": userID,
		"id":      userSourceID,
	}
	us := userSourceSrv.UserSourceDAO.Find(cond)
	if len(userSourceID) != len(us) {
		return false
	}

	return true
}
