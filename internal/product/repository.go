package product

type ProductRepository interface {
	GetAll() []Product
	Store(p Product) Product
	FindByID(id int) *Product
}

type ProductRepo struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepo{}
}

func (p *ProductRepo) GetAll() []Product {
	return ProductList
}

func (p *ProductRepo) Store(prod Product) Product {
	prod.ID = len(ProductList) + 1
	ProductList = append(ProductList, prod)
	return prod
}

func (p *ProductRepo) FindByID(id int) *Product {
	for _, prod := range ProductList {
		if prod.ID == id {
			return &prod
		}
	}
	return nil
}
