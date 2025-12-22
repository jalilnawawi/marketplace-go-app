package controller_impl

import (
	"net/http"
	"strconv"

	"github.com/jalilnawawi/marketplace-app/controller"
	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/dto"
	"github.com/jalilnawawi/marketplace-app/model/dto/order_dto"
	"github.com/jalilnawawi/marketplace-app/service"
	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) controller.OrderController {
	return &OrderControllerImpl{OrderService: orderService}
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Create a new order with name, price, and description
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order  body      order_dto.CreateOrderRequest  true  "Create Order Request"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/order [post]
func (controller *OrderControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderCreateRequest := order_dto.CreateOrderRequest{}
	helper.ReadFromRequestBody(request, &orderCreateRequest)

	orderResponse, err := controller.OrderService.Create(request.Context(), orderCreateRequest)
	if err != nil {
		apiResponse := dto.ApiResponse{
			Code:   400,
			Status: "bad request",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, apiResponse)
	} else {
		apiResponse := dto.ApiResponse{
			Code:   201,
			Status: "created",
			Data:   orderResponse,
		}
		helper.WriteToResponseBody(writer, apiResponse)
	}
}

// GetAllOrders godoc
// @Summary      Get all orders
// @Description  Retrieve a list of all orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/order [get]
func (controller *OrderControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderResponses := controller.OrderService.FindAll(request.Context())
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "ok",
		Data:   orderResponses,
	}
	helper.WriteToResponseBody(writer, apiResponse)
}

// GetOrderById godoc
// @Summary      Get order by ID
// @Description  Retrieve a order by its ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        orderId   path      int  true  "Order ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/order/{orderId} [get]
func (controller *OrderControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.FindById(request.Context(), int64(id))
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "ok",
		Data:   orderResponse,
	}
	helper.WriteToResponseBody(writer, apiResponse)
}

// UpdateOrder godoc
// @Summary      Update a order
// @Description  Update a order's information by its ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        orderId  path      int  true  "Order ID"
// @Param        order    body      order_dto.UpdateOrderRequest  true  "Update Order Request"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/order/{orderId} [put]
func (controller *OrderControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateOrderRequest := order_dto.UpdateOrderRequest{}
	helper.ReadFromRequestBody(request, &updateOrderRequest)

	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	updateResponse := controller.OrderService.Update(request.Context(), updateOrderRequest, int64(id))

	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "ok",
		Data:   updateResponse,
	}
	helper.WriteToResponseBody(writer, apiResponse)

}

// DeleteOrder godoc
// @Summary      Delete a order
// @Description  Delete a order by its ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        orderId  path      int  true  "Order ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/order/{orderId} [delete]
func (controller *OrderControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	controller.OrderService.Delete(request.Context(), int64(id))

	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "ok",
		Data:   "order deleted successfully",
	}
	helper.WriteToResponseBody(writer, apiResponse)
}
