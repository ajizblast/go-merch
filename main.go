// package name
package main

// import package
import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Struct
type Todos struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

//struct digunakan untuk mendefinisikan sebuah kata kunci ketika mau membuat object atau array of object
//diatas ada sebuah struct bernama todos yang berisi Id bertipe data string, Tittle ini juga string, IsDone ini boolean
//properti json yang dipakai nanti adalah yg berada di "", misal json:"id"

// Dummy Data

// diatas ada array of object dan berisi 2 object

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		// c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().WriteHeader(http.StatusOK)
		return c.String(http.StatusOK, "Hello World!")
	})

	fmt.Println("Server running on localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
