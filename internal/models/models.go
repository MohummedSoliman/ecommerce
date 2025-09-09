// Package models define the type of database connection values.
package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Models is wrapper for all models.
type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventroyLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

type Order struct {
	ID            int       `json:"id"`
	WidgetID      int       `json:"widget_id"`
	TransactionID int       `json:"transaction_id"`
	StatusID      int       `json:"status_id"`
	Quantity      int       `json:"quantity"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type Status struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TransactionStatus struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Transaction struct {
	ID                  int       `json:"id"`
	Amount              int       `json:"amount"`
	Currency            string    `json:"currency"`
	LastFour            string    `json:"last_four"`
	BankReturnCode      string    `json:"bank_return_code"`
	TransactionStatusID int       `json:"transaction_status_id"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (m *DBModel) GetWedgit(id int) (Widget, error) {
	ctx, cencel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cencel()

	var widget Widget

	query := `SELECT id, name, description, inventory_level, price,
				coalesce(image, ''), created_at, updated_at FROM widgets WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(&widget.ID, &widget.Name, &widget.Description, &widget.InventroyLevel,
		&widget.Price, &widget.Image, &widget.CreatedAt, &widget.UpdatedAt)
	if err != nil {
		return widget, err
	}

	if err = row.Err(); err != nil {
		return widget, err
	}

	return widget, nil
}

func (m *DBModel) InsertTransaction(t Transaction) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO transactions (amount, currency, last_four, bank_return_code, transaction_status_id)
			  VALUES (?, ?, ?, ?, ?)`
	result, err := m.DB.ExecContext(ctx, stmt, t.Amount, t.Currency, t.LastFour, t.BankReturnCode, t.TransactionStatusID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *DBModel) InsertOrder(o Order) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO order (widget_id, transaction_id, status_id, quantity, amount)
			 VALUES (?, ?, ?, ?, ?)`
	result, err := m.DB.ExecContext(ctx, stmt, o.WidgetID, o.TransactionID, o.StatusID, o.Quantity, o.Amount)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
