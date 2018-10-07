package comm

type Response struct {
	Data       interface{}
	Status     ResponseStatus `json:"status"`
	TotalCount int            `json:"tcnt,omitempty"`
}

type ResponseStatus struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *ResponseStatus) Error() string {
	return s.Msg
}

var statusOK = ResponseStatus{Code: 0, Msg: ""}

func ResponseOk(data interface{}) Response {
	return ResponseOkMsg(data, ``)
}

func ResponseOkMsg(data interface{}, msg string) (resp Response) {
	resp.Status = statusOK
	resp.Status.Msg = msg
	resp.Data = data
	return
}

func ResponseErr(data interface{}, code int, msg string) (resp Response) {
	resp.Status.Code = code
	resp.Status.Msg = msg
	resp.Data = data
	return
}
