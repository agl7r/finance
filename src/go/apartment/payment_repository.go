package apartment

import (
	"database/sql"
	"fmt"
	"github.com/agl7r/finance/config"
	"github.com/bojanz/currency"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

type PaymentRepository struct {
}

func getDB() *sql.DB {
	dsn := config.Config.Dsn
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func NewPaymentRepository() *PaymentRepository {
	return new(PaymentRepository)
}

func (r *PaymentRepository) FindAll() (CommunalPayments, error) {
	var payments CommunalPayments

	db := getDB()

	res, err := db.Query("SELECT * FROM payments")
	if res != nil {
		defer res.Close()
	}
	if err != nil {
		return payments, err
	}

	for res.Next() {
		var payment CommunalPayment

		var monthId string
		var typeId int
		var amount string

		err := res.Scan(&monthId, &typeId, &amount)
		if err != nil {
			return payments, err
		}

		payment.Month = &Month{Id: monthId}
		payment.Type = GetTypes().GetById(typeId)
		payment.Amount, _ = currency.NewAmount(amount, "RUB")

		payments = append(payments, &payment)
	}

	return payments, nil
}

func (r *PaymentRepository) FindByMonth(m *Month) (CommunalPayments, error) {
	var payments CommunalPayments

	allPayments, err := r.FindAll()
	if err != nil {
		return payments, err
	}
	for _, payment := range allPayments {
		if payment.Month.Id == m.Id {
			payments = append(payments, payment)
		}
	}

	return payments, nil
}

func (r *PaymentRepository) FindByYear(year int) (CommunalPayments, error) {
	payments := CommunalPayments{}

	allPayments, _ := r.FindAll()
	for _, payment := range allPayments {
		parts := strings.Split(payment.Month.Id, "-")
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

func (r *PaymentRepository) Save(payment *CommunalPayment) error {
	db := getDB()

	amount := payment.Amount.Number()
	_, err := db.Query(fmt.Sprintf("INSERT INTO payments VALUES ('%s', %d, %s) ON DUPLICATE KEY UPDATE amount = %s",
		payment.Month.Id, payment.Type.Id, amount, amount))
	if err != nil {
		return err
	}

	return nil
}
