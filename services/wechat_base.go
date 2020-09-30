package services

import (
	"encoding/json"
	"fmt"
	"lee_wechat/models"
	"lee_wechat/utils"
)

func GetAccessToken() {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxe7c90d3859e08158&secret=3a82cca3518cca0853d8cb2962d237c1"
	responseByte := utils.HttpGet(url, "")
	var accessTokenResponse models.AccessTokenResponse
	err := json.Unmarshal(responseByte, &accessTokenResponse)
	if err != nil {
		panic(err)
	}
	if accessTokenResponse.ErrCode != 0 {
		panic(accessTokenResponse.ErrMsg)
	}
	fmt.Println(accessTokenResponse.AccessToken)
}
