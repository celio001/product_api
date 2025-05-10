package product

import (
    "context"
    "database/sql"
    "testing"

    "github.com/stretchr/testify/assert"
    _ "github.com/mattn/go-sqlite3" // Driver SQLite para testes
)

func setupTestDB() (*sql.DB, error) {
    // Cria um banco de dados SQLite em memória
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        return nil, err
    }

    // Cria a tabela de produtos para os testes
    createTableQuery := `
    CREATE TABLE product (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        product_name TEXT NOT NULL,
        price REAL NOT NULL,
        description TEXT,
        stock INTEGER NOT NULL DEFAULT 0,
        category TEXT
    );`
    _, err = db.Exec(createTableQuery)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func TestNewProduct(t *testing.T) {
    // Configura o banco de dados de teste
    db, err := setupTestDB()
    assert.NoError(t, err)
    defer db.Close()

    // Cria uma instância do repositório
    repo := &repository{sqlDb: db}

    // Dados do produto a ser criado
    addProduct := &Product{
        Name:        "Test Product",
        Price:       10.0,
        Description: "Test Description",
        Stock:       100,
        Category:    "Test Category",
    }

    // Executa o método de criação
    ctx := context.Background()
    id, err := repo.NewProduct(ctx, addProduct)

    // Verifica se não houve erro e se o ID retornado é válido
    assert.NoError(t, err)
    assert.NotZero(t, id)
}