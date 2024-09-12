package types

type UserRegisterReq struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

type UserLoginReq struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

type UserInfoResp struct {
	Id       int64  `json:"id"`
	UserName string `json:"name"`
	CreateAt int64  `json:"create_at"`
}

type TokenData struct {
	User        UserInfoResp `json:"user"`
	AccessToken string       `json:"access_token"`
}
