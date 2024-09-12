package types

type CreateTaskReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}

type ListTaskReq struct {
	Limit int `json:"limit"`
	Start int `json:"start"`
}

type ListTaskResp struct {
	Uid       int64  `json:"uid"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	View      int64  `json:"view"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type ShowTaskReq struct {
	Id int64 `json:"id"`
}

type UpdateTaskReq struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}
type DeleteTaskReq struct {
	Id int64 `json:"id"`
}

type SearchTaskReq struct {
	Info string `json:"info"`
}
