package response

import "github.com/gofiber/fiber/v2"

// Success is
type Success struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewSuccess is
func NewSuccess(c *fiber.Ctx, code int, m string, d interface{}) error {
	data := Success{
		Error:   false,
		Message: m,
		Data:    d,
	}
	return c.Status(code).JSON(data)

}
