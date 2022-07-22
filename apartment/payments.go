package apartment

import "github.com/bojanz/currency"

type CommunalPayment struct {
	Month  *Month          `json:"month"`
	Type   *Type           `json:"type"`
	Amount currency.Amount `json:"amount"`
}

type CommunalPayments []*CommunalPayment

func (payments *CommunalPayments) GetTotal() currency.Amount {
	var total currency.Amount

	for _, p := range *payments {
		total, _ = total.Add(p.Amount)
	}

	return total
}

func NewCommunalPayment(month *Month, paymentType *Type, amount currency.Amount) *CommunalPayment {
	payment := new(CommunalPayment)
	payment.Month = month
	payment.Type = paymentType
	payment.Amount = amount
	return payment
}
