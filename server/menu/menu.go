package menu

import (
	"github.com/jmoiron/sqlx"
)

type MenuItem struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type MenuFacade struct {
	db *sqlx.DB
}

func NewMenuFacade(db *sqlx.DB) *MenuFacade {
	return &MenuFacade{db: db}
}

func (m MenuFacade) GetAllMenuItems() ([]*MenuItem, error) {
	var res []*MenuItem
	err := m.db.Select(&res, "SELECT ID, NAME, PRICE FROM MENU_ITEMS")
	return res, err
}

func (m MenuFacade) GetMenuItems(ids []int) ([]*MenuItem, error) {
	query, args, _ := sqlx.In("SELECT ID, NAME, PRICE FROM MENU_ITEMS WHERE ID in (?)", ids)
	q := m.db.Rebind(query)
	var res []*MenuItem
	err := m.db.Select(&res, q, args...)
	return res, err
}
