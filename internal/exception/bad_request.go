package exception

import "encoding/json"

type ErrBadRequest Err

func BadRequest(err interface{}) ErrBadRequest {
	var msg string
	switch err.(type) {
	case string:
		msg = err.(string)
	case error:
		msg = err.(error).Error()
	}

	return ErrBadRequest{
		ErrorType: TypeErrorBadRequest,
		ErrorCode: 400,
		Message:   msg,
	}
}

func (e ErrBadRequest) Error() string {
	var msg string
	if IsHttpError {
		payload, _ := json.Marshal(e)
		msg = string(payload)
	}

	return msg
}
