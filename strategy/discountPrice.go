package strategy

import "github.com/Zhu-Ying-Zhang/bitgin-payment/utils"

type DiscountPrice struct {
	platformPoint int
	discountRatio float64
}

func NewDiscountPrice(platformPoint int, discountRatio float64) *DiscountPrice {
	return &DiscountPrice{platformPoint: platformPoint, discountRatio: discountRatio}
}

func (discount *DiscountPrice) GetTotalPay(price int) int {
	return int(utils.CalculatePlatformPointDiscountPrice(float64(price), discount.platformPoint, discount.discountRatio))
}