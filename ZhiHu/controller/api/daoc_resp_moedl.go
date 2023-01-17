package api

import "Little_Project/ZhiHu/tools"

type _ResponsePostList struct {
	Code    int           `json:"code"`    // 业务响应状态码
	Message string        `json:"message"` // 提示信息
	Data    []*tools.Resp `json:"data"`    // 数据
}
