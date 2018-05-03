package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/yosssi/ace"
)

func main() {
	e := echo.New()
	e.GET("/test", aceTest)
	e.GET("/", helloWorld)
	e.GET("/readthread", readThread)
	e.POST("/postshoot", postMessage)
	e.Logger.Fatal(e.Start(":1323"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "HELLO WORLD"+strconv.Itoa(100))
}

func postMessage(c echo.Context) error {
	return c.String(http.StatusOK, "HE")
}

func readThread(c echo.Context) error {
	return c.String(http.StatusBadRequest, "bad")
}

func aceTest(c echo.Context) error {
	tpl, err := ace.Load("template/base", "inner", nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err := tpl.Execute(c.Response(), map[string]string{"Msg": "Hello Ace"}); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "test")
}
