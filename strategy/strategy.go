package strategy

type Strategy interface {
    GetTotalPay(price int) int
}