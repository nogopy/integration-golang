package handler

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tuanbieber/integration-golang/internal/model"
	"github.com/tuanbieber/integration-golang/internal/service"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PhoneHandler struct {
	PhoneService service.PhoneServiceInterface
}

func NewPhoneHandler(phoneService service.PhoneServiceInterface) *PhoneHandler {
	return &PhoneHandler{PhoneService: phoneService}
}

func (h *PhoneHandler) CreateOnePhone(c echo.Context) error {
	var phone model.Phone

	if err := c.Bind(&phone); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &Response{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := h.PhoneService.CreateOnePhone(&phone); err != nil {
		httpStatus := http.StatusInternalServerError
		if reflect.TypeOf(err).String() == `validator.ValidationErrors` {
			httpStatus = http.StatusBadRequest
		}

		return c.JSON(httpStatus, &Response{
			Status:  httpStatus,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]int{
			"id": phone.ID,
		},
	})
}

func (h *PhoneHandler) GetOnePhoneById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	res, err := h.PhoneService.GetOnePhoneById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if res == nil {
		return c.JSON(http.StatusNotFound, &Response{
			Status:  http.StatusNotFound,
			Message: "phone is not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    res,
	})
}

func (h *PhoneHandler) GetAllPhone(c echo.Context) error {
	res, err := h.PhoneService.GetAllPhone()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    res,
	})
}
