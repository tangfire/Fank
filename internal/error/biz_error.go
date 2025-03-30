package biz_err

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (b *Err) Error() string {
	return b.Msg
}

func New(code int, msg ...string) *Err {
	message := ""

	if len(msg) <= 0 {
		message = GetMessage(code)
	} else {
		message = msg[0]
	}

	return &Err{Code: code, Msg: message}
}
