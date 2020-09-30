package controller

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"lee_wechat/models"
	"sort"
	"strings"
	"time"
)

const (
	token = "siaslfs"
)

// @Summary test
// @Tags test
// @Accept plain
// @Produce json
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string	"fail"
// @Router /wechat/receive/msg [post]
func ReceiveMsgFromWechat(c echo.Context) error {
	return c.JSON(200, "ok")
}

// @Summary test
// @Tags test
// @Accept plain
// @Produce json
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string	"fail"
// @Router /wechat/receive/check [get]
func CheckSignature(c echo.Context) error {
	if c.Request().Method == "GET" {
		signature := c.FormValue("signature")
		timestamp := c.FormValue("timestamp")
		nonce := c.FormValue("nonce")
		echostr := c.FormValue("echostr")

		validateBool := validateUrl(signature, timestamp, nonce)
		if validateBool {
			return c.String(200, echostr)
		}
		return c.String(400, "验证失败")
	} else if c.Request().Method == "POST" {
		var msgResponse models.MsgResponse
		if err := c.Bind(&msgResponse); err != nil {
			return c.String(400, "参数解析失败")
		}
		newMsgResponse := models.MsgResponse{
			XMLName:      xml.Name{},
			ToUserName:   msgResponse.FromUserName,
			FromUserName: msgResponse.ToUserName,
			CreateTime:   time.Duration(time.Now().Unix()),
			MsgType:      msgResponse.MsgType,
			Content:      fmt.Sprintf("公众号回复信息：%s", msgResponse.Content),
		}
		return c.XML(200, newMsgResponse)
	}
	return c.JSON(200, "ok")
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func validateUrl(signature, timestamp, nonce string) bool {
	signatureGen := makeSignature(timestamp, nonce)
	if signatureGen != signature {
		return false
	}
	return true
}
