package response

type ResponseMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponsePageMsg struct {
	ResponseMsg
	PageSize uint `json:"pageSize"`
	PageNum  uint `json:"pageNum"`
	total    uint `json:"totalPage"`
}

func SuccessMsg(data interface{}) *ResponseMsg {
	msg := &ResponseMsg{
		Code: 0,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func SuccessMsgForPages(data interface{}, pageSize uint, pageNum uint, total uint) *ResponsePageMsg {
	msg := &ResponsePageMsg{
		ResponseMsg: ResponseMsg{
			Code: 0,
			Msg:  "SUCCESS",
			Data: data,
		},
		PageSize: pageSize,
		PageNum:  pageNum,
		total:    total,
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
