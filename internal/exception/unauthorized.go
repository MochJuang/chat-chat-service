package exception

import "encoding/json"

type ErrUnauthorized Err

func Unauthorized(err interface{}) ErrUnauthorized {
	var msg string
	switch err.(type) {
	case string:
		msg = err.(string)
	case error:
		msg = err.(error).Error()
	}

	return ErrUnauthorized{
		ErrorType: TypeErrorUnauthorized,
		ErrorCode: 401,
		Message:   msg,
	}
}

func (e ErrUnauthorized) Error() string {
	var msg string
	if IsHttpError {
		payload, _ := json.Marshal(e)
		msg = string(payload)
	}

	return msg
}
