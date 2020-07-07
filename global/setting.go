package global

import (
	"github.com/Blackoutta/blog-service/pkg/logger"
	"github.com/Blackoutta/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
