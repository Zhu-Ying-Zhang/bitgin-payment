package strategy

import "github.com/zhu-ying-zhang/bitgin-payment/utils"

type VipDiscountPrice struct {
	vip int
	platformPoint int
	discountRatio float64
}

func NewVipDiscountPrice(vip int, platformPoint int, discountRatio float64) *VipDiscountPrice {
	return &VipDiscountPrice{ vip: vip, platformPoint: platformPoint, discountRatio: discountRatio }
}

func (vipDiscount *VipDiscountPrice) GetTotalPay(price int) int {

	discountPrice := utils.CalculatePlatformPointDiscountPrice(float64(price), vipDiscount.platformPoint, vipDiscount.discountRatio)

	extraDiscountPrice := discountPrice *  extraDiscountRatio(vipDiscount.vip, vipDiscount.platformPoint)

	return int(extraDiscountPrice)
}

func extraDiscountRatio(vip int, platformPoint int) float64 {
	if vip > 0 && platformPoint >= 100 {
		return 0.9
	} else {
		return 1.0
	}
}