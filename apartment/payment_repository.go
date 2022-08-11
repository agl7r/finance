package apartment

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const jsonFilePath = "var/apartment/payments.json"

type PaymentRepository struct {
}

func NewPaymentRepository() *PaymentRepository {
	return new(PaymentRepository)
}

func (r *PaymentRepository) FindAll() (CommunalPayments, error) {
	var payments CommunalPayments

	content, _ := ioutil.ReadFile(jsonFilePath)
	_ = json.Unmarshal(content, &payments)

	return payments, nil
}

func (r *PaymentRepository) FindByMonth(m *Month) (CommunalPayments, error) {
	var payments CommunalPayments

	allPayments, _ := r.FindAll()
	for _, payment := range allPayments {
		if payment.Month.Id == m.Id {
			payments = append(payments, payment)
		}
	}

	return payments, nil
}

func (r *PaymentRepository) FindByYear(year int) (CommunalPayments, error) {
	var payments CommunalPayments

	allPayments, _ := r.FindAll()
	for _, payment := range allPayments {
		parts := strings.Split(payment.Month.Id,"-")
		paymentYear, _ := strconv.Atoi(parts[0])
		if paymentYear < 100 {
			paymentYear += 2000
		}
		if paymentYear == year {
			payments = append(payments, payment)
		}
	}

	return payments, nil
}

func (r *PaymentRepository) FindByMonthAndType(m *Month, t Type) (CommunalPayments, error) {
	var payments CommunalPayments

	allPayments, _ := r.FindAll()
	for _, payment := range allPayments {
		if payment.Month.Id == m.Id && payment.Type.Id == t.Id {
			payments = append(payments, payment)
		}
	}

	return payments, nil
}

func (r *PaymentRepository) Save(p *CommunalPayment) error {
	var payments []*CommunalPayment

	content, _ := ioutil.ReadFile(jsonFilePath)
	_ = json.Unmarshal(content, &payments)

	found := false
	for idx, payment := range payments {
		if payment.Month.Id == p.Month.Id && payment.Type.Id == p.Type.Id {
			payments[idx] = p
			found = true
		}
	}

	if found == false {
		payments = append(payments, p)
	}

	sort.Slice(payments, func(i, j int) bool {
		return payments[i].Month.Id < payments[j].Month.Id
	})

	newContent, _ := json.MarshalIndent(payments, "", " ")

	_ = ioutil.WriteFile(jsonFilePath, newContent, 0644)

	return nil
}
