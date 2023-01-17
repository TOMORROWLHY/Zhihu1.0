package tools

//向前端反馈信息

type Resp struct {
	Code    int         `json:"code"`           // 业务响应状态码
	Message interface{} `json:"message"`        // 提示信息
	Data    interface{} `json:"data,omitempty"` // 数据
}

type respTemplate struct {
	Status int    `json:"Status"`
	info   string `json:"Info"`
}

var OK = respTemplate{
	Status: 200,
	info:   "success",
}

var ParamError = respTemplate{
	Status: 300,
	info:   "params error",
}

var InternalError = respTemplate{
	Status: 500,
	info:   "internal error",
}
