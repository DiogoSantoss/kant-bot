package main

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/courses", getCourses)
	e.GET("/course/:id", getCourse)

	e.Logger.Fatal(e.Start(":8081"))
}

func getCourses(c echo.Context) error {

	// With MEIC-A id
	endpoint := "https://fenix.tecnico.ulisboa.pt/api/fenix/v1/degrees/2761663971475/courses"
	
	resp, err := http.Get(endpoint)
	if err != nil {
		c.Echo().Logger.Error(err)
		return err
	}

	encodedJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Echo().Logger.Error(err)
		return err
	}

	return c.JSONBlob(http.StatusOK, encodedJSON)
}

func getCourse(c echo.Context) error {

	id := c.Param("id")
	endpoint := "https://fenix.tecnico.ulisboa.pt/api/fenix/v1/courses/"
	endpoint += id

	resp, err := http.Get(endpoint)
	if err != nil {
		c.Echo().Logger.Error(err)
		return err
	}

	encodedJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Echo().Logger.Error(err)
		return err
	}

	return c.JSONBlob(http.StatusOK, encodedJSON)
}
