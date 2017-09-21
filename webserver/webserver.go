package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

type Cat  struct {
	name	string	`json:"name"`
	Type	string	`json:"type"`
}

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the web side!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\nand his type is: %s\n", catName, catType))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string {
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string {
		"error": "you need to let us know if you want json or string data",
	})

}

func addCats(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Print("Failed reading the request body for addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarshalling in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "we got you cat")
}

func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()

	e.GET("/", yallo)
	e.GET("/cats/:data", getCats)

	e.POST("/cats", addCats)

	e.Start(":8000")
}