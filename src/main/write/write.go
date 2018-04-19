package write

import (
	"net/http"

	"github.com/labstack/echo"
)

// WriteStatus は入力APIに投げるJSON
type writeStatus struct {
	Theme string `json:"theme"`
	Query string `json:"query"`
}

// Write は入力処理を行うJSON
func Write(context echo.Context) error {
	ws := new(writeStatus)
	if err := context.Bind(ws); err != nil {
		// エラー処理
		return err
	}
	return context.JSON(http.StatusCreated, ws)

	// 書き込みデータをスレッドサーバに飛ばす
}
