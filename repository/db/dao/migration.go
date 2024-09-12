package dao

import "demo/repository/db/model"

func migrate() {
	err := _db.Set("gorm:table_settings", "charset=utf8&parseTime=True&loc=Local").AutoMigrate(&model.TaskModel{}, &model.UserModel{})
	if err != nil {
		panic(err)
	}
}
