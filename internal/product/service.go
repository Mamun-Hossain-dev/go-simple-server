package product

import "fmt"

type ProductService interface {
	StoreProduct(p Product) Product
	GetAllProducts() []Product
	GetProductByID(id int) (*Product, error)
}

type productService struct {
	repo ProductRepository
}

func NewProductService(r ProductRepository) ProductService {
	return &productService{
		repo: r,
	}
}

func (s *productService) StoreProduct(p Product) Product {
	return s.repo.Store(p)
}

func (s *productService) GetAllProducts() []Product {
	return s.repo.GetAll()
}

func (s *productService) GetProductByID(id int) (*Product, error) {
	prod := s.repo.FindByID(id)
	if prod == nil {
		return nil, fmt.Errorf("product with ID %d not found", id)
	}
	return prod, nil
}
