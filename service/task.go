package service

import (
	"context"
	"demo/pkg/ctl"
	"demo/pkg/utils"
	"demo/repository/db/dao"
	"demo/repository/db/model"
	"demo/types"
	"sync"
	"time"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

// GetTaskSrv 通过Once完成小型结构体的实例化，而不需要显式地写在其他地方
func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (*TaskSrv) CreateTask(c context.Context, req *types.CreateTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	task := &model.TaskModel{
		Uid:       u.Id,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	err = dao.NewTaskDao(c).CreateTask(task)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (t *TaskSrv) ListTask(c context.Context, req *types.ListTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	tasks, total, err := dao.NewTaskDao(c).ListTask(req.Start, req.Limit, u.Id)
	tRespList := make([]*types.ListTaskResp, 0)
	for _, task := range tasks {
		tRespList = append(tRespList, &types.ListTaskResp{
			Uid:       task.Uid,
			Title:     task.Title,
			Content:   task.Content,
			Status:    task.Status,
			View:      task.View(),
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
			CreatedAt: task.CreatedAt.Unix(),
		})
	}
	return ctl.RespList(tRespList, total), nil
}

func (t *TaskSrv) ShowTask(c context.Context, req *types.ShowTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	task, err := dao.NewTaskDao(c).FindTaskById(req.Id, u.Id)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	tResp := &types.ListTaskResp{
		Uid:       task.Uid,
		Title:     task.Title,
		Content:   task.Content,
		Status:    task.Status,
		View:      task.View(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
		CreatedAt: task.CreatedAt.Unix(),
	}
	return ctl.RespSuccessWithData(tResp), nil
}

func (t *TaskSrv) DeleteTask(c context.Context, req *types.DeleteTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	err = dao.NewTaskDao(c).DeleteTaskByIdAndUid(req.Id, u.Id)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (t *TaskSrv) UpdateTask(c context.Context, req *types.UpdateTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	err = dao.NewTaskDao(c).UpdateTask(req, u.Id)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (t *TaskSrv) SearchTask(c context.Context, req *types.SearchTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	tasks, err := dao.NewTaskDao(c).SearchTask(req.Info, u.Id)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	tRespList := make([]*types.ListTaskResp, 0)
	for _, task := range tasks {
		tRespList = append(tRespList, &types.ListTaskResp{
			Uid:       task.Uid,
			Title:     task.Title,
			Content:   task.Content,
			Status:    task.Status,
			View:      task.View(),
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
			CreatedAt: task.CreatedAt.Unix(),
		})
	}
	return ctl.RespSuccessWithData(tRespList), nil
}
