package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var hoge = rand.Int()

func main() {
	e := echo.New()
	e.GET("/", helloWorld)
	e.GET("/readthread", readThread)
	e.POST("/postshoot", postMessage)
	e.Logger.Fatal(e.Start(":1323"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "HELLO WORLD"+strconv.Itoa(hoge))
}

func postMessage(c echo.Context) error {
	return c.String(http.StatusOK, "HE")
}

func readThread(c echo.Context) error {
	return c.String(http.StatusBadRequest, "bad")
}
