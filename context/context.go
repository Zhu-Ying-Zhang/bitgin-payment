package context

import "github.com/zhu-ying-zhang/bitgin-payment/strategy"

// TODO factory
type Context struct {
    strategy strategy.Strategy
}

func NewContext(str strategy.Strategy) *Context {
    return &Context{
        strategy: str,
    }
}

func (c *Context) ExecuteStrategy(price int) int {
    return c.strategy.GetTotalPay(price)
}