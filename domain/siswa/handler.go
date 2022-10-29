package siswa

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/dot-golang/util/logger"
	"github.com/jadahbakar/dot-golang/util/response"
)

type Handler struct {
	service Service
}

func NewHandler(route fiber.Router, s Service) {
	handler := &Handler{service: s}
	route.Get("/health", handler.GetHealth)
	route.Post("/siswa", handler.Post)
	route.Get("/siswa", handler.GetAllSiswa)
	route.Get("/siswa/:nis", handler.GetByNis)
	route.Put("/siswa/:nis", handler.Update)
	route.Delete("/siswa/:nis", handler.Delete)

}

func (h *Handler) GetHealth(c *fiber.Ctx) error {
	return response.NewSuccess(c, fiber.StatusOK, "healthy", nil)
}

func (h *Handler) Post(c *fiber.Ctx) error {
	s := &Siswa{}
	if err := c.BodyParser(s); err != nil {
		logger.Errorf("Error On Body Parser -> ", err)
		return response.BadRequest(c, err.Error())
	}
	id, err := h.service.PostSiswa(s)
	if err != nil {
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Post", id)
}

func (h *Handler) GetAllSiswa(c *fiber.Ctx) error {
	listData, err := h.service.FindAllSiswa()
	if err != nil {
		logger.Error(err)
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Find All Siswa", listData)
}

func (h *Handler) GetByNis(c *fiber.Ctx) error {
	id := c.Params("nis")
	data, err := h.service.FindByNIS(id)
	if err != nil {
		logger.Error(err)
		return response.NoRowsInResultSet(c, err.Error())
	}
	return response.NewSuccess(c, fiber.StatusOK, "Search By Nis", data)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("nis")
	s := new(Siswa)
	if err := c.BodyParser(s); err != nil {
		return response.BadRequest(c, err.Error())
	}

	RowsAffected, err := h.service.UpdateSiswa(id, s)
	if err != nil {
		logger.Error(err)
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Update Siswa", RowsAffected)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("nis")

	RowsAffected, err := h.service.DeleteSiswa(id)
	if err != nil {
		logger.Error(err)
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Delete Siswa", RowsAffected)
}
