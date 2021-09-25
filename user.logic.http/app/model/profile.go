package model

type ProfileApiDetailResp struct {
	UpdatedAt   int64  `json:"updated_at"`
	Uid         int64  `json:"uid"`         // 用户ID
	Name        string `json:"name"`        // 用户昵称
	Passport    string `json:"passport"`    // 用户账号
	Gender      int    `json:"gender"`      // 性别
	Birth       int64  `json:"birth"`       // 生日
	Portrait    string `json:"portrait"`    // 头像
	Hometown    string `json:"hometown"`    // 家乡
	Description string `json:"description"` // 简介
}
