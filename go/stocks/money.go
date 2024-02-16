package stocks

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) Money {
	return Money{amount: amount, currency: currency}
}

func (m Money) Times(multiplier int) Money {
	return Money{currency: m.currency, amount: float64(multiplier) * m.amount}
}

func (m Money) Divide(divisor int) Money {
	return Money{currency: m.currency, amount: m.amount / float64(divisor)}
}
