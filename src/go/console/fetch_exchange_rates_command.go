package console

import (
	"encoding/xml"
	"fmt"
	"github.com/bojanz/currency"
	"golang.org/x/net/html/charset"
	"net/http"
	"strings"
)

type FetchExchangeRatesCommand struct{}

func (c *FetchExchangeRatesCommand) Names() []string {
	return []string{"fetch-exchange-rates", "rates"}
}

func (c *FetchExchangeRatesCommand) Execute(args []string) error {
	url := "https://www.cbr.ru/scripts/XML_daily.asp"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("запрос завершился ошибкой: %d", resp.StatusCode)
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
		return fmt.Errorf("не удалось распарсить ответ: %s", err)
	}

	var rates []Rate
	for _, r := range decoded.Valutes {
		a, _ := currency.NewAmount(strings.Replace(r.Value, ",", ".", 1), "RUB")
		a, _ = a.Div(fmt.Sprintf("%d", r.Nominal))
		rates = append(rates, Rate{r.CharCode, r.Name, a})
	}

	fmt.Printf("Курсы валют на %s:\n", decoded.Date)
	for _, rate := range rates {
		fmt.Printf("%s (%s): %s\n", rate.Name, rate.CharCode, rate.Value)
	}

	return nil
}

type Rate struct {
	CharCode string
	Name     string
	Value currency.Amount
}
