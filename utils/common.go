package utils

func CalculateVipPrice(price float64, discount float64) float64 {
	return price * discount
}

func CalculatePlatformPointDiscountPrice(price float64, platformPoint int, discountRatio float64) float64 {
	return price - (float64(platformPoint) * discountRatio)
}