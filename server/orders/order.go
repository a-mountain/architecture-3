package orders

import (
	"github.com/jmoiron/sqlx"
	"github.com/nickname038/architecture-3/menu"
)

type Order struct {
	Id          int   `json:"id"`
	TableNumber int   `json:"table_number"`
	MenuItemId  []int `json:"menu_item_id"`
}

type OrderPrice struct {
	Price        float64 `json:"price"`
	PriceNoTax   float64 `json:"price_no_tax"`
	RecommendTip float64 `json:"recommend_tip"`
}

type CreateOrderRequest struct {
	TableNumber int
	MenuItemIds []int
}

type OrderFacade struct {
	tipFactor  float64
	taxFactor  float64
	db         *sqlx.DB
	menuFacade *menu.MenuFacade
}

func NewOrderFacade(db *sqlx.DB, facade *menu.MenuFacade, tipFactor float64, taxFactor float64) *OrderFacade {
	return &OrderFacade{db: db, taxFactor: taxFactor, tipFactor: tipFactor, menuFacade: facade}
}

func (f OrderFacade) CreateOrder(request CreateOrderRequest) (OrderPrice, error) {
	tx := f.db.MustBegin()
	for _, id := range request.MenuItemIds {
		number := request.TableNumber
		tx.MustExec("INSERT INTO ORDERS (TABLE_NUMBER, MENU_ITEM_ID) VALUES ($1, $2)", number, id)
	}
	err := tx.Commit()
	if err != nil {
		return OrderPrice{}, err
	}
	menuItems, err := f.menuFacade.GetMenuItems(request.MenuItemIds)
	return f.calcPrice(menuItems), err
}

func (f OrderFacade) calcPrice(menuItems []*menu.MenuItem) OrderPrice {
	priceNoTax := sumMenuItemsPrice(menuItems)
	realPrice := priceNoTax + (priceNoTax * f.taxFactor)
	tip := priceNoTax * f.tipFactor
	return OrderPrice{
		PriceNoTax:   priceNoTax,
		RecommendTip: tip,
		Price:        realPrice,
	}
}

func sumMenuItemsPrice(menuItems []*menu.MenuItem) float64 {
	var price = 0.0
	for _, item := range menuItems {
		price += item.Price
	}
	return price
}
