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

// CreateSeller godoc
// @Summary Create a new seller
// @Description Create a new seller with name, category, and description
// @Tags sellers
// @Accept json
// @Produce json
// @Param seller body seller_dto.CreateSellerRequest true "Create Seller Request"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seller [post]
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

// GetAllSeller godoc
// @Summary Get all sellers
// @Description Retrieve a list of all sellers
// @Tags sellers
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seller [get]
func (controller *SellerControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sellerResponses := controller.SellerService.FindAll(request.Context())
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
		Data:   sellerResponses,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

// GetByIdSeller godoc
// @Summary Get seller by ID
// @Description Retrieve a seller by its ID
// @Tags sellers
// @Accept json
// @Produce json
// @Param sellerId path int true "Seller ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seller/{sellerId} [get]
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

// UpdateSeller godoc
// @Summary Update a seller
// @Description Update a seller's information by its ID
// @Tags sellers
// @Accept json
// @Produce json
// @Param sellerId path int true "Seller ID"
// @Param seller body seller_dto.CreateSellerRequest true "Update Seller Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seller/{sellerId} [put]
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

// DeleteSeller godoc
// @Summary Delete a seller
// @Description Delete a seller by its ID
// @Tags sellers
// @Accept json
// @Produce json
// @Param sellerId path int true "Seller ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seller/{sellerId} [delete]
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
