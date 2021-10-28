package types


// Представляет собой денежную сумму в минимальных единицах (копейки и т.д.)
type Money int64

// Представляет код валюты
type Currency string

// Category представляет собой категорию, в которой был совершен платеж
type Category string

// Коды валют

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// Представляет номер карты
type PAN string

// Представляет информацию о платежной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}



type Payment struct {
	ID int 
	Amount Money
	Category Category
}


type PaymentSource struct {
	Type    string // card
	Number  string // номер вида '5058 хххх хххх 8888' 
	Balance Money // баланс в дирамах
}



