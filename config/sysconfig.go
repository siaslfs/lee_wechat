package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type SysConfig struct {
	QiNiu struct {
		AccessKey string
		SecretKey string
		Bucket    string
		BaseUrl   string
	}
	Default struct {
		Avatar     string
		DataBase   string
		Omnipotent string //是否使用万能验证码
	}
	Aliyun struct {
		AccessKeyId     string
		AccessKeySecret string
		TemplateCode    string
		SignName        string
	}
	Wechat struct {
		Appid  string
		Secret string
	}
}

func (m *SysConfig) LoadConfig() (service SysConfig) {
	_, err := toml.DecodeFile("appsetting.toml", &m)
	if err != nil {
		log.Fatal(err)
	}
	return *m
}
