package product

type Product struct {
	ID          int     `json:"id"`
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    Description string  `json:"description"` // Descrição do produto
    Stock       int     `json:"stock"`       // Quantidade em estoque
    Category    string  `json:"category"`    // Categoria do produto
    CreatedAt   string  `json:"created_at"`  // Data de criação
    UpdatedAt   string  `json:"updated_at"`  // Data de última atualização
}