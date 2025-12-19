package controller_impl

import (
	"net/http"
	"strconv"

	"github.com/jalilnawawi/marketplace-app/controller"
	"github.com/jalilnawawi/marketplace-app/helper"
	"github.com/jalilnawawi/marketplace-app/model/dto"
	"github.com/jalilnawawi/marketplace-app/model/dto/product_dto"
	"github.com/jalilnawawi/marketplace-app/service"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) controller.ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product with name, price, and description
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      product_dto.CreateProductRequest  true  "Create Product Request"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/product [post]
func (p *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := product_dto.CreateProductRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)

	productCreateResponse, err := p.ProductService.Create(request.Context(), productCreateRequest)
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
			Data:   productCreateResponse,
		}
		helper.WriteToResponseBody(writer, apiResponse)
	}
}

// GetAllProducts godoc
// @Summary      Get all products
// @Description  Retrieve a list of all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/product [get]
func (p *ProductControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponses := p.ProductService.FindAll(request.Context())
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

// GetProductById godoc
// @Summary      Get product by ID
// @Description  Retrieve a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        productId   path      int  true  "Product ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/product/{productId} [get]
func (p *ProductControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := p.ProductService.FindById(request.Context(), int64(id))
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update a product's information by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        productId  path      int  true  "Product ID"
// @Param        product    body      product_dto.UpdateProductRequest  true  "Update Product Request"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/product/{productId} [put]
func (p *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := product_dto.UpdateProductRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := p.ProductService.Update(request.Context(), productUpdateRequest, int64(id))

	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writer, apiResponse)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        productId  path      int  true  "Product ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/product/{productId} [delete]
func (p *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	p.ProductService.Delete(request.Context(), int64(id))
	apiResponse := dto.ApiResponse{
		Code:   200,
		Status: "success",
	}
	helper.WriteToResponseBody(writer, apiResponse)
}
