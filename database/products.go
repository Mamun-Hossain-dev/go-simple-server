package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var ProductList []Product

type Response struct {
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

type ProductResponse struct {
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

func LoadProducts() {

	prod1 := Product{
		ID:          1,
		Title:       "Banana",
		Description: "I love banana. It has high rich photasiam",
		Price:       125.223,
		ImgUrl:      "https://images.unsplash.com/photo-1571771894821-ce9b6c11b08e?ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YmFuYW5hfGVufDB8fDB8fHww&fm=jpg&q=60&w=3000",
	}

	prod2 := Product{
		ID:          2,
		Title:       "Banana",
		Description: "I love banana. It has high rich photasiam",
		Price:       125.223,
		ImgUrl:      "https://images.unsplash.com/photo-1571771894821-ce9b6c11b08e?ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YmFuYW5hfGVufDB8fDB8fHww&fm=jpg&q=60&w=3000",
	}

	prod3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "I love banana. It has high rich photasiam",
		Price:       125.223,
		ImgUrl:      "https://images.unsplash.com/photo-1571771894821-ce9b6c11b08e?ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YmFuYW5hfGVufDB8fDB8fHww&fm=jpg&q=60&w=3000",
	}

	// append this products to ProductList
	ProductList = append(ProductList, prod1)
	ProductList = append(ProductList, prod2)
	ProductList = append(ProductList, prod3)
}
