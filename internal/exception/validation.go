package exception

import "encoding/json"

type ErrValidation Err

func Validation(err interface{}) ErrValidation {
	var msg string
	switch err.(type) {
	case string:
		msg = err.(string)
	case error:
		msg = err.(error).Error()
	}

	return ErrValidation{
		ErrorType: TypeErrorValidation,
		ErrorCode: 400,
		Message:   msg,
	}
}

func (e ErrValidation) Error() string {
	var msg string
	if IsHttpError {
		payload, _ := json.Marshal(e)
		msg = string(payload)
	}

	return msg
}
