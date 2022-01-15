package response

type ResponseMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponsePageMsg struct {
	ResponseMsg
	PageSize  int `json:"pageSize"`
	PageNum   int `json:"pageNum"`
	TotalPage int `json:"totalPage"`
}

func SuccessMsg(data interface{}) *ResponseMsg {
	msg := &ResponseMsg{
		Code: 0,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func SuccessMsgForPages(data interface{}, pageSize int, pageNum int, total int) *ResponsePageMsg {
	msg := &ResponsePageMsg{
		ResponseMsg: ResponseMsg{
			Code: 0,
			Msg:  "SUCCESS",
			Data: data,
		},
		PageSize:  pageSize,
		PageNum:   pageNum,
		TotalPage: total,
	}
	return msg
}

func FailMsg(msg string) *ResponseMsg {
	msgObj := &ResponseMsg{
		Code: -1,
		Msg:  msg,
	}
	return msgObj
}
