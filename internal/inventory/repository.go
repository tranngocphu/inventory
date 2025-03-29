package inventory

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Define a database object
type Repository struct {
	db *sqlx.DB
}

// Create database connection
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// Define items table in the databse
type Item struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Stock int    `db:"stock"`
}

// Add new row to items table
func (r *Repository) CreateItem(item Item) error {
	_, err := r.db.Exec("INSERT INTO items (name, stock) VALUES ($1, $2)", item.Name, item.Stock)
	return err
}

// Retrieve an item from the items table
func (r *Repository) GetItem(id int) (Item, error) {
	var item Item
	err := r.db.Get(&item, "SELECT * FROM items WHERE id = $1", id)
	return item, err
}
