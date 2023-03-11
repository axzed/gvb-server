package flag

import (
	"gvb-server/global"
	"gvb-server/model"
)

func Makemigrations() {
	var err error
	if err = global.DB.SetupJoinTable(&model.UserModel{}, "CollectsModels", &model.UserCollectModel{}); err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	if err = global.DB.SetupJoinTable(&model.MenuModel{}, "Banners", &model.MenuBannerModel{}); err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.BannerModel{},
			//&model.TagModel{},
			//&model.MessageModel{},
			&model.AdvertModel{},
			//&model.UserModel{},
			//&model.CommentModel{},
			//&model.ArticleModel{},
			&model.MenuModel{},
			&model.MenuBannerModel{},
			&model.FeedBackModel{},
			//&model.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
