package bayar

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/dot-golang/util/logger"
	"github.com/jadahbakar/dot-golang/util/response"
)

type Handler struct {
	service Service
}

func NewHandler(route fiber.Router, s Service) {
	handler := &Handler{service: s}
	route.Post("/bayar/", handler.Post)
	route.Put("/bayar/", handler.Update)
	route.Delete("/bayar/:nis/:idbayar", handler.Delete)
	route.Get("/bayar/:nis", handler.GetById)
	route.Get("/bayar", handler.GetAllBayar)
}

func (h *Handler) Post(c *fiber.Ctx) error {
	b := &Bayar{}
	if err := c.BodyParser(b); err != nil {
		logger.Errorf("Error On Body Parser -> ", err)
		return response.BadRequest(c, err.Error())
	}
	id, err := h.service.PostBayar(b)
	if err != nil {
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Post", id)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	b := new(Bayar)
	if err := c.BodyParser(b); err != nil {
		return response.BadRequest(c, err.Error())
	}

	RowsAffected, err := h.service.UpdateBayar(b)
	if err != nil {
		logger.Error(err)
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Update Bayar", RowsAffected)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	nis := c.Params("nis")
	id := c.Params("idbayar")
	idBayar, _ := strconv.ParseInt(id, 10, 64)

	RowsAffected, err := h.service.DeleteBayar(nis, idBayar)
	if err != nil {
		logger.Error(err)
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Delete Bayar", RowsAffected)
}

func (h *Handler) GetById(c *fiber.Ctx) error {
	id := c.Params("nis")
	data, err := h.service.FindById(id)
	if err != nil {
		logger.Error(err)
		return response.NoRowsInResultSet(c, err.Error())
	}
	return response.NewSuccess(c, fiber.StatusOK, "Search By Nis", data)
}

func (h *Handler) GetAllBayar(c *fiber.Ctx) error {
	listData, err := h.service.FindAllBayar()
	if err != nil {
		logger.Error(err)
		return response.HandleErrors(c, err)
	}
	return response.NewSuccess(c, fiber.StatusOK, "Find All Bayar", listData)

}
