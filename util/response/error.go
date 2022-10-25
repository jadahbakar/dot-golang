package response

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
)

// errors response
var (
	ErrBadRequest          = errors.New("Bad Request, something wrong on your request") // 400
	ErrFailedGenerateToken = errors.New("Failed Generate Token")                        // 500
	ErrInternalServer      = errors.New("Internal Server Error")                        // 500
	ErrDuplicate           = errors.New("The desired shortcode is already in use")      // 409
)

// Error is
type Error struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Stack   interface{} `json:"stack"`
}

// NewError is
func NewError(f *fiber.Ctx, code int, m string, s interface{}) error {
	data := Error{
		Error:   true,
		Message: m,
		Stack:   s,
	}
	return f.Status(code).JSON(data)
}

// BadRequest is | 400
func BadRequest(f *fiber.Ctx, s interface{}) error {
	return NewError(f, fiber.StatusBadRequest, ErrBadRequest.Error(), nil)
}

// HandleErrors is
func HandleErrors(c *fiber.Ctx, e error) error {
	// 400
	if errors.Is(e, ErrBadRequest) {
		return NewError(c, fiber.StatusBadRequest, ErrBadRequest.Error(), nil)
	}

	// duplicate 409
	if errors.Is(e, e.(*pgconn.PgError)) {
		return NewError(c, fiber.StatusConflict, ErrDuplicate.Error(), nil)
	}

	// internal server error
	return NewError(c, fiber.StatusInternalServerError, ErrInternalServer.Error(), nil)
}

// Fiber Default Error Handler for Config
var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError
	message := ErrInternalServer.Error()
	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	// Return statuscode with error message
	data := Error{true, message, nil}
	return c.Status(code).JSON(data)
}
