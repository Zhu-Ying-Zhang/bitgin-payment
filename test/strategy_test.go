package test

import (
	"testing"

	"github.com/zhu-ying-zhang/bitgin-payment/context"
	"github.com/zhu-ying-zhang/bitgin-payment/data"
	"github.com/zhu-ying-zhang/bitgin-payment/strategy"
)

type User struct {
	name          string
	vip           int
	platformPoint int
}

func NewUser(name string, vip int, platformPoint int) *User {
	return &User{name: name, vip: vip, platformPoint: platformPoint}
}

func Test(t *testing.T) {
	t.Run("Normal price test:", NormalPriceStrategyTest)
	t.Run("Discount price test", DiscountPriceStrategyTest)
	t.Run("VIP price test:", VIPPriceStrategyTest)
	t.Run("VIP Discount price test", VIPDiscountPriceStrategyTest)
}

func NormalPriceStrategyTest(t *testing.T) {
	price := 1000
	expectedPrice := 1000
	c := context.NewContext(strategy.NewNormalPrice())
	result := c.ExecuteStrategy(price)
	
	if result != expectedPrice {
		t.Error("The result of normal price should equal to expectedPrice price!")
	}
}

func DiscountPriceStrategyTest(t *testing.T) {
	j := NewUser("Joe", 0, 100)
	price := 1000
	spendPlatformPoint := 100
	expectedPrice := 900
	if spendPlatformPoint > j.platformPoint {
		t.Error("The spending point is higher than user's platform point!")
	}
	c := context.NewContext(strategy.NewDiscountPrice(spendPlatformPoint, data.DiscountRatio))
	result := c.ExecuteStrategy(price)
	
	if result != expectedPrice {
		t.Error("The result of discount price should equal to expected price!")
	}
}

func VIPPriceStrategyTest(t *testing.T) {
	j := NewUser("Joe", 0, 50)
	b := NewUser("Brian", 2, 1000)
	price := 1000
	firstUserExpectedResult := 1000
	c := context.NewContext(strategy.NewVipPrice(data.VipRate[j.vip]))
	result := c.ExecuteStrategy(price)

	if result != firstUserExpectedResult {
		t.Error("The result of discount price of First user should equal to expected price!")
	}
	secondUserExpectedResult := 900
	c = context.NewContext(strategy.NewVipPrice(data.VipRate[b.vip]))
	result = c.ExecuteStrategy(price)

	if result != secondUserExpectedResult {
		t.Error("The result of discount price of Second user should equal to expected price!")
	}
}

func VIPDiscountPriceStrategyTest(t *testing.T) {
	j := NewUser("Joe", 0, 100)
	b := NewUser("Brian", 2, 100)
	i := NewUser("Ivy", 2, 50)
	price := 1000

	firstUserExpectedResult := 900
	firstUserSpendPlatformPoint := 100
	if firstUserSpendPlatformPoint > j.platformPoint {
		t.Error("The spending point is higher than user's platform point!")
	}
	c := context.NewContext(strategy.NewVipDiscountPrice(j.vip, firstUserSpendPlatformPoint, data.DiscountRatio))
	result := c.ExecuteStrategy(price)
	
	if result != firstUserExpectedResult {
		t.Error("The result of discount price of First user should equal to expected price!")
	}

	secondUserExpectedResult := 810
	secondUserSpendPlatformPoint := 100
	if secondUserSpendPlatformPoint > j.platformPoint {
		t.Error("The spending point is higher than user's platform point!")
	}
	c = context.NewContext(strategy.NewVipDiscountPrice(b.vip, secondUserSpendPlatformPoint, data.DiscountRatio))
	result = c.ExecuteStrategy(price)
	
	if result != secondUserExpectedResult {
		t.Error("The result of discount price of Second user should equal to expected price!")
	}

	thirdUserExpectedResult := 950
	thirdUserSpendPlatformPoint := 50
	if thirdUserSpendPlatformPoint > j.platformPoint {
		t.Error("The spending point is higher than user's platform point!")
	}
	c = context.NewContext(strategy.NewVipDiscountPrice(i.vip, thirdUserSpendPlatformPoint, data.DiscountRatio))
	result = c.ExecuteStrategy(price)
	
	if result != thirdUserExpectedResult {
		t.Error("The result of discount price of Third user should equal to expected price!")
	}
}
