package controller_impl

import (
	"net/http"
	"strconv"

	"github.com/jalilnawawi/marketplace-app/controller"
	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/dto"
	"github.com/jalilnawawi/marketplace-app/model/dto/seller_dto"
	"github.com/jalilnawawi/marketplace-app/service"
	"github.com/julienschmidt/httprouter"
)

type SellerControllerImpl struct {
	SellerService service.SellerService
}

func NewSellerController(sellerService service.SellerService) controller.SellerController {
	return &SellerControllerImpl{SellerService: sellerService}
}

// CreateItem godoc
// @Summary Create a new seller
// @Description Create a new seller with name, category, and description
// @Tags sellers
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sellers [post]
func (controller *SellerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sellerCreateRequest := seller_dto.CreateSellerRequest{}
	helper.ReadFromRequestBody(request, &sellerCreateRequest)

	sellerCreateResponse := controller.SellerService.Create(request.Context(), sellerCreateRequest)
	apiResponse := dto.ApiResponse{
		Code:   201,
		Status: "created",
		Data:   sellerCreateResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *SellerControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sellerResponses := controller.SellerService.FindAll(request.Context())
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
		Data:   sellerResponses,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *SellerControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sellerId := params.ByName("sellerId")
	id, err := strconv.Atoi(sellerId)
	helper.PanicIfError(err)

	sellerResponse := controller.SellerService.FindById(request.Context(), int64(id))
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
		Data:   sellerResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *SellerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sellerUpdateRequest := seller_dto.CreateSellerRequest{}
	helper.ReadFromRequestBody(request, &sellerUpdateRequest)

	sellerId := params.ByName("sellerId")
	id, err := strconv.Atoi(sellerId)
	helper.PanicIfError(err)

	sellerResponse := controller.SellerService.Update(request.Context(), sellerUpdateRequest, int64(id))
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
		Data:   sellerResponse,
	}
	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *SellerControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sellerId := params.ByName("sellerId")
	id, err := strconv.Atoi(sellerId)
	helper.PanicIfError(err)

	controller.SellerService.Delete(request.Context(), int64(id))
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
	}
	helper.WriteToResponseBody(writer, apiResponse)
}
