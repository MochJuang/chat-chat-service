package exception

import "encoding/json"

type ErrInternal Err

func Internal(err interface{}) ErrInternal {
	var msg string
	switch err.(type) {
	case string:
		msg = err.(string)
	case error:
		msg = err.(error).Error()
	}

	return ErrInternal{
		ErrorCode: 500,
		ErrorType: TypeErrorInternal,
		Message:   msg,
	}
}

func (e ErrInternal) Error() string {
	var msg string
	payload, _ := json.Marshal(e)
	msg = string(payload)

	return msg
}
