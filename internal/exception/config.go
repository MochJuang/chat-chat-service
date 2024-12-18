package exception

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

const IsHttpError = true
const (
	TypeErrorNotFound     = "NotFound"
	TypeErrorBadRequest   = "BadRequest"
	TypeErrorValidation   = "Validation"
	TypeErrorUnauthorized = "Unauthorized"
	TypeErrorInternal     = "Internal"
)

var (
	NotFoundError     = NotFound("")
	BadRequestError   = BadRequest("")
	InternalError     = Internal("")
	UnauthorizedError = Unauthorized("")
)

type Err struct {
	ErrorType string
	ErrorCode int
	Message   string
}

func (e Err) Error() string {
	var msg string
	if IsHttpError {
		payload, _ := json.Marshal(e)
		msg = string(payload)
	}

	return msg
}

func Convert(err error) (Err, error) {
	var e Err
	newErr := json.Unmarshal([]byte(err.Error()), &e)
	if newErr != nil {
		return e, newErr
	}

	return e, nil
}

func HandleHttpErrorFiber(c *fiber.Ctx, err error) error {
	var msg, _err = Convert(err)

	if _err != nil {
		log.Println(err.Error())
		log.Println("Error parsing error message middleware")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error"})
	}

	log.Println(msg)

	switch msg.ErrorType {
	case TypeErrorValidation:
		var validationsErr []map[string]interface{}
		_ = json.Unmarshal([]byte(msg.Message), &validationsErr)
		return c.Status(msg.ErrorCode).JSON(fiber.Map{"errors": validationsErr})
	default:
		return c.Status(msg.ErrorCode).JSON(fiber.Map{"message": msg.Message})
	}
}
