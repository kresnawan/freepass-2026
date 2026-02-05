package api

type Error struct {
	Msg      string
	HttpCode int
}

type ResponseMessage struct {
	Succ int    `json:"success"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func MakeError(msg string, httpcode int) *Error {
	return &Error{
		Msg:      msg,
		HttpCode: httpcode,
	}
}

func MakeResponse(succ int, msg string, data any) ResponseMessage {
	return ResponseMessage{
		Succ: succ,
		Msg:  msg,
		Data: data,
	}
}
