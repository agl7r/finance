package usecases

import (
	"encoding/xml"
	"fmt"
	"github.com/bojanz/currency"
	"golang.org/x/net/html/charset"
	"net/http"
	"strings"
	"time"
)

type FetchExchangeRates struct {

}

func (u *FetchExchangeRates) Execute(date time.Time) (Rates, error) {
	rates := Rates{}

	url := "https://www.cbr.ru/scripts/XML_daily.asp"

	resp, err := http.Get(url)
	if err != nil {
		return rates, fmt.Errorf("не удалось выполнить запрос: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return rates, fmt.Errorf("запрос завершился ошибкой: %d", resp.StatusCode)
	}

	decoded := struct {
		XMLName xml.Name `xml:"ValCurs"`
		Date    string   `xml:"Date,attr"`
		Valutes []struct {
			XMLName  xml.Name `xml:"Valute"`
			NumCode  string   `xml:"NumCode"`
			CharCode string   `xml:"CharCode"`
			Nominal  int      `xml:"Nominal"`
			Name     string   `xml:"Name"`
			Value    string   `xml:"Value"`
		} `xml:"Valute"`
	}{}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&decoded)
	if err != nil {
		return rates, fmt.Errorf("не удалось распарсить ответ: %s", err)
	}

	rates.Date = decoded.Date
	for _, r := range decoded.Valutes {
		a, _ := currency.NewAmount(strings.Replace(r.Value, ",", ".", 1), "RUB")
		a, _ = a.Div(fmt.Sprintf("%d", r.Nominal))
		rates.Rates = append(rates.Rates, Rate{r.CharCode, r.Name, a})
	}

	return rates, nil
}

type Rate struct {
	CharCode string
	Name     string
	Value currency.Amount
}

type Rates struct {
	Date string
	Rates []Rate
}

func (rates *Rates) ByCode(code string) *Rate {
	for _, r := range rates.Rates {
		if r.CharCode == code {
			return &r
		}
	}
	return nil
}
