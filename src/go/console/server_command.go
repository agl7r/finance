package console

import (
	"fmt"
	"github.com/agl7r/finance/apartment"
	"github.com/agl7r/finance/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ServerCommand struct {}

func (c *ServerCommand) Names() []string {
	return []string{"server"}
}

func getUtilityBillsEndpoint(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Не указан год")
		return
	}

	response := struct {
		Bills apartment.CommunalPayments `json:"bills"`
	}{}

	repository := apartment.NewPaymentRepository()
	bills, err := repository.FindByYear(year)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, "Не удалось получить счета")
		return
	} else {
		response.Bills = bills
	}

	c.JSON(http.StatusOK, response)
}

func (c *ServerCommand) Execute(args []string) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	v1.GET("/utility-bills/:year", getUtilityBillsEndpoint)

	router.Run(fmt.Sprintf(":%d", config.Config.Port))

	return nil
}
