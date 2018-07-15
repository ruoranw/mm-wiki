package controllers

import (
	"mm-wiki/app/models"
	"strings"
)

type ConfigController struct {
	BaseController
}

func (this *ConfigController) Global() {

	configs, err := models.ConfigModel.GetConfigs()
	if err != nil {
		this.ErrorLog("获取全局配置失败: "+err.Error())
		this.ViewError("邮件服务器不存在", "/system/main/index")
	}

	var configValue = map[string]string{}
	for _, config := range configs{
		if config["key"] == "auto_follow_doc_open" && config["value"] != "1" {
			config["value"] = "0"
		}
		if config["key"] == "send_email_open" && config["value"] != "1" {
			config["value"] = "0"
		}
		if config["key"] == "sso_open" && config["value"] != "1" {
			config["value"] = "0"
		}
		configValue[config["key"]] = config["value"]
	}

	this.Data["configValue"] = configValue
	this.viewLayout("config/form", "default")
}

func (this *ConfigController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/email/list")
	}
	mainTitle := this.GetString("main_title", "")
	mainDescription := strings.TrimSpace(this.GetString("main_description", ""))
	autoFollowDocOpen := strings.TrimSpace(this.GetString("auto_follow_doc_open", "0"))
	sendEmailOpen := strings.TrimSpace(this.GetString("send_email_open", "0"))
	ssoOpen := strings.TrimSpace(this.GetString("sso_open", "0"))

	_, err := models.ConfigModel.UpdateByKey("main_title", mainTitle)
	if err != nil {
		this.ErrorLog("修改配置 main_title  失败: "+err.Error())
		this.jsonError("主页标题配置失败！")
	}

	_, err = models.ConfigModel.UpdateByKey("main_description", mainDescription)
	if err != nil {
		this.ErrorLog("修改配置 main_description  失败: "+err.Error())
		this.jsonError("主页描述配置失败！")
	}

	_, err = models.ConfigModel.UpdateByKey("auto_follow_doc_open", autoFollowDocOpen)
	if err != nil {
		this.ErrorLog("修改配置 auto_follow_doc_open  失败: "+err.Error())
		this.jsonError("开启自动关注配置失败！")
	}

	_, err = models.ConfigModel.UpdateByKey("send_email_open", sendEmailOpen)
	if err != nil {
		this.ErrorLog("修改配置 send_email_open  失败: "+err.Error())
		this.jsonError("开启邮件通知配置失败！")
	}

	_, err = models.ConfigModel.UpdateByKey("sso_open", ssoOpen)
	if err != nil {
		this.ErrorLog("修改配置 sso_open  失败: "+err.Error())
		this.jsonError("开启统一登录配置失败！")
	}

	this.InfoLog("修改全局配置成功")
	this.jsonSuccess("修改全局配置成功", nil, "/system/config/global")
}