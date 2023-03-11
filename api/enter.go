package api

import "gvb-server/api/settings"

// ApiGroupApp 实例化api组
var ApiGroupApp = new(ApiGroup)

// ApiGroup api组
type ApiGroup struct {
	SettingsApi settings.SettingsApi
}
