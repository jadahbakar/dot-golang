package response

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// errors response
var (
	ErrBadRequest          = errors.New("Bad Request, something wrong on your request")           // 400
	ErrInvalidCredential   = errors.New("Please provide valid credentials")                       // 401
	ErrUnauthorized        = errors.New("Need Authorizetion Credential")                          // 401
	ErrNotFound            = errors.New("Not Found, your request data not found in our database") // 404
	ErrNotFoundRoute       = errors.New("Not Found, URL you want is not in this application")     // 404
	ErrMethodNotAllowed    = errors.New("Method Not Allowed")                                     // 405
	ErrUnprocessableEntity = errors.New("Validation Failed")                                      // 422
	ErrFailedGenerateToken = errors.New("Failed Generate Token")                                  // 500
	ErrInternalServer      = errors.New("Internal Server Error")                                  // 500
	ErrNoRows              = errors.New("no rows in result set")                                  // 503

	// validation message
	minMsg      = "value must be at least"
	requiredMsg = "value must be required"
	emailMsg    = "value must be a valid email"
	defaultMsg  = "value must be validate"
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

// Unauthorized is | 401
func Unauthorized(f *fiber.Ctx, s interface{}) error {
	return NewError(f, fiber.StatusUnauthorized, ErrUnauthorized.Error(), nil)
}

// UnprocessableEntityVar is
func UnprocessableEntityVar(f *fiber.Ctx, m map[string]interface{}) error {
	return NewError(f, fiber.StatusUnprocessableEntity, ErrUnprocessableEntity.Error(), m)
}

// NoRowsInResultSet is
func NoRowsInResultSet(f *fiber.Ctx, s interface{}) error {
	return NewError(f, fiber.StatusServiceUnavailable, ErrNoRows.Error(), nil)
}

// HandleErrors is
func HandleErrors(c *fiber.Ctx, e error) error {

	if IsSQLError(e) {
		code := GetSQLErrorCode(e)
		if code == "23505" {
			return NewError(c, fiber.StatusBadRequest, ErrBadRequest.Error(), nil)
		}
	}

	// login failed | 401
	if errors.Is(e, ErrInvalidCredential) {
		return NewError(c, fiber.StatusUnauthorized, ErrInvalidCredential.Error(), nil)
	}

	// login failed or failed generate token | 500
	if errors.Is(e, ErrFailedGenerateToken) {
		return NewError(c, fiber.StatusInternalServerError, ErrFailedGenerateToken.Error(), nil)
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

var CustomErrorHandler = func(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	// Send custom error page
	// err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(code).SendString("Internal Server Error")
	}
	// Return from handler
	return nil
}
