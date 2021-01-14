package services

import (
	"goproject/datamodels"
	"goproject/repositories"
)

type ProductService interface {
	InsertProduct(product *datamodels.Product) (int64, error)
	DeleteProduct(productId int64) bool
	UpdateProduct(product *datamodels.Product) error
	GetProductByID(productId int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
}

type ProductServiceManager struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &ProductServiceManager{productRepository: productRepository}
}

func (p *ProductServiceManager) InsertProduct(product *datamodels.Product) (int64, error) {
	//panic("implement me")
	return p.productRepository.Insert(product)
}

func (p *ProductServiceManager) DeleteProduct(productId int64) bool {
	//panic("implement me")
	return p.productRepository.Delete(productId)
}

func (p *ProductServiceManager) UpdateProduct(product *datamodels.Product) error {
	//panic("implement me")
	return p.productRepository.Update(product)
}

func (p *ProductServiceManager) GetProductByID(productId int64) (*datamodels.Product, error) {
	//panic("implement me")
	return p.productRepository.SelectByKey(productId)
}

func (p *ProductServiceManager) GetAllProduct() ([]*datamodels.Product, error) {
	//panic("implement me")
	return p.productRepository.SelectAll()
}
