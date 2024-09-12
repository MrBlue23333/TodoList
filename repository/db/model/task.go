package model

import (
	"demo/repository/cache"
	"github.com/spf13/cast"
	"time"
)

type TaskModel struct {
	Id        int64      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	Uid       int64      `gorm:"column:uid"`
	Title     string     `gorm:"column:title"`
	Status    int        `gorm:"column:status"`
	Content   string     `gorm:"column:content"`
	StartTime int64      `gorm:"column:start_time"`
	EndTime   int64      `gorm:"column:end_time"`
}

func (t *TaskModel) TableName() string {
	return "task"
}
func (t *TaskModel) View() int64 {
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(t.Id)).Result()
	return int64(cast.ToInt(countStr))
}

func (t *TaskModel) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(t.Id))
}
