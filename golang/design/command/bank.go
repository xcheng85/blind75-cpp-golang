package command

type IBankCommand interface {
	// like ngrx reducer
	Execute(balance float64) float64
}

type DepositeCommand struct {
	amount float64
}

func (c *DepositeCommand) Execute(balance float64) float64 {
	return c.amount + balance
}

type WidthdrawCommand struct {
	amount float64
}

func (c *WidthdrawCommand) Execute(balance float64) float64 {
	return balance - c.amount
}
