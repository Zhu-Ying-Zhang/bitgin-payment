package strategy

import "github.com/zhu-ying-zhang/bitgin-payment/utils"

type VipPrice struct {
	discount float64
}

func NewVipPrice(discount float64) *VipPrice {
	return &VipPrice{ discount: discount }
}

func (vip *VipPrice) GetTotalPay(price int) int {
	return int(utils.CalculateVipPrice(float64(price), vip.discount))
}