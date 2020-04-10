package models

type Book struct {
	ID    int     `db:"id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
}
