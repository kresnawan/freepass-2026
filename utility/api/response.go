package api

type ResponseMessage struct {
	Msg          string `json:"msg"`
	ErrMsg       string `json:"err_msg"`
	RowsAffected int64  `json:"rows_affected"`
	LastInsertId int64  `json:"last_insert_id"`
}

func MakeResponse(msg string, errmsg string, rows ...int64) ResponseMessage {

	var rA int64 = 0
	var lI int64 = 0

	if len(rows) > 0 {
		rA = rows[0]
	}

	if len(rows) > 1 {
		lI = rows[1]
	}

	return ResponseMessage{
		Msg:          msg,
		ErrMsg:       errmsg,
		RowsAffected: rA,
		LastInsertId: lI,
	}
}
