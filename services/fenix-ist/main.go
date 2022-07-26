package main

import (
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

/*
File structure
https://www.codementor.io/@packt/how-to-set-up-a-project-in-echo-mpo39w5zu
*/

func main() {

	e := echo.New()

	e.GET("/teste", func(c echo.Context) error {

		endpoint := "https://fenix.tecnico.ulisboa.pt/api/fenix/v1/courses/1610612925989?lang=en-US"
		client := &http.Client{Timeout: 10 * time.Second}

		req, _ := http.NewRequest("GET", endpoint, nil)
		res, _ := client.Do(req)

		// Close the connection to reuse it
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		return c.String(http.StatusOK, string(body))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
