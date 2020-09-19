package delivery

import (
	"net/http"
	"rest_api/domain"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type customerOrderHandler struct {
	orderUcase domain.CustomerOrderUsecase
}

func NewCustomerOrderHandler(e *echo.Echo, u domain.CustomerOrderUsecase) {
	handler := &customerOrderHandler{
		orderUcase: u,
	}

	e.GET("/order-details/:name", handler.Order)
}

func (h *customerOrderHandler) Order(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.orderUcase.GetOrdersByCustomer(ctx, c.Param("name"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
