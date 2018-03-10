package factory

import (
	"errors"
	"fmt"
)

// Payment methods type
type PaymentType int

const (
	Cash PaymentType = iota + 1
	CreditCard
	DebitCard
)

//PaymentMethod defines a way of paying in the shop. This factory method returns
//objects that implements this interface
type PaymentMethod interface {
	Pay(amount float32) string
	GetBalance() float32
}

// CreatePaymentMethod returns a pointer to a PaymentMethod object or an error
func GetPaymentMethod(m PaymentType) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

//-----------------------------------------------------------------------------
// Cash
//-----------------------------------------------------------------------------
type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

func (c *CashPM) GetBalance() float32 {
	return 100.05
}

//-----------------------------------------------------------------------------
// CreditCard
//-----------------------------------------------------------------------------
type CreditCardPM struct{}

func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using credit card\n", amount)
}

func (c *CreditCardPM) GetBalance() float32 {
	return 100.05
}

//-----------------------------------------------------------------------------
// DebitCard
//-----------------------------------------------------------------------------
type DebitCardPM struct{}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card\n", amount)
}

func (c *DebitCardPM) GetBalance() float32 {
	return 100.05
}
