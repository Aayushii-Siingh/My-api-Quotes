package main

import "github.com/labstack/echo/v4"
import "net/http"
import "time"
import "math/rand"
import "os"


type quote struct {
	Title string
	Author string
}

//var quotes []quote = make([]quote, 0)
var quotes []quote = []quote {
	{"Learn to lead", "Sai vidya campus"},
	{"Eat strawberries", "everyday"},
	{"i am quote", "3"},
	{"blahaha blahah blahhh", "4 here"},
}
func main() {

	rand.Seed(time.Now().Unix())

	api := echo.New()
	api.GET("/quotes", getQuotes)
	api.GET("/quotes/random", getRandomQuotes)
	//endpoint---handler
	//api.GET("/", hello)
	//api.POST("/", hello)
	//api.PUT("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}

	api.Start(":" + port)
}

func getQuotes(c echo.Context) error{
	return c.JSON(http.StatusOK, quotes)
}

func getRandomQuotes(c echo.Context) error {
	
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])
}

func hello(c echo.Context) error {
	return c.String(200, "Hello Aayushi")
}