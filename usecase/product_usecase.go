package usecase

import (
	"github.com/celio001/product_api.git/model"
	"github.com/celio001/product_api.git/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProduct() ([]model.Product, error) {
	return pu.repository.GetProduct() 
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error){
	productId, err := pu.repository.CreateProduct(product)
	if err != nil{
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id_product int) (*model.Product, error){

	product, err := pu.repository.GetProductById(id_product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUseCase) DeleteProcuctById(id_product int) (*model.Product, error){

	product, err := pu.repository.DeleteProcuctById(id_product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(productUpdate model.Product) (*model.Product, error){
	var product *model.Product
	
	product, err := pu.repository.UpdateProduct(productUpdate)

	if err != nil {
		return nil, err
	}

	return product, nil
}