package dao

import (
	"context"
	"demo/pkg/utils"
	"demo/repository/db/model"
	"demo/types"
	"gorm.io/gorm"
)

type TaskDao struct {
	db *gorm.DB
}

func NewTaskDao(c context.Context) *TaskDao {
	if c == nil {
		c = context.Background()
	}
	return &TaskDao{NewDbClient(c)}
}

func (t *TaskDao) CreateTask(task *model.TaskModel) error {
	return t.db.Create(task).Error
}

func (t *TaskDao) ListTask(start, limit int, uId int64) (tasks []*model.TaskModel, total int64, err error) {
	err = t.db.Model(&model.TaskModel{}).
		Where("uid = ?", uId).
		Count(&total).
		Error
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}

	err = t.db.Model(&model.TaskModel{}).
		Where("uid = ?", uId).
		Limit(limit).
		Offset((start - 1) * limit).
		Find(&tasks).
		Error
	return
}
func (t *TaskDao) FindTaskById(id, uid int64) (task *model.TaskModel, err error) {
	err = t.db.Model(&model.TaskModel{}).
		Where("id = ? AND uid = ?", id, uid).
		First(&task).
		Error
	return
}

func (t *TaskDao) DeleteTaskByIdAndUid(id, uid int64) error {
	return t.db.Model(&model.TaskModel{}).
		Where("id = ? AND uid = ?", id, uid).
		Delete(&model.TaskModel{}).
		Error
}

func (t *TaskDao) UpdateTask(req *types.UpdateTaskReq, uid int64) error {
	var task model.TaskModel
	err := t.db.Model(&model.TaskModel{}).
		Where("id = ? AND uid = ?", req.Id, uid).
		First(&task).
		Error
	if err != nil {
		return err
	}
	if req.Status != 0 {
		task.Status = req.Status
	}
	if req.Content != "" {
		task.Content = req.Content
	}
	if req.Title != "" {
		task.Title = req.Title
	}
	return t.db.Save(&task).Error
}

func (t *TaskDao) SearchTask(info string, uid int64) (tasks []*model.TaskModel, err error) {
	err = t.db.Model(&model.TaskModel{}).
		Where("uid = ? AND (title LIKE ? OR content LIKE ?)", uid, "%"+info+"%", "%"+info+"%").
		Find(&tasks).
		Error
	return
}
