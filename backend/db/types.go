package db

type Order struct {
	Id         int
	ProviderId string
	Name       string
	Price      float32
	OrderDate  datefmt
}

type Provider struct {
	Id       int
	Name     string
	LastSync datefmt
	Username string
	Password string
}
