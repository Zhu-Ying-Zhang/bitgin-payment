package strategy

type NormalPrice struct {}

func NewNormalPrice() *NormalPrice {
	return &NormalPrice{}
}

func (normal *NormalPrice) GetTotalPay(price int) int {
	return price
}