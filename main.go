package main

import "github.com/labstack/echo/v4"
import "net/http"
import "math/rand"
import "time"
import "os"

type quote struct {

	Title string 
	Author string
}

//var quotes []quote = make([]quote,0)
var quotes []quote = []quote{
	{"Learn to Lead","Sai Vidya Campus"},
	{"quote 2","author 2"},
	{"quote 3","author 3"},
	{"quote 4","author 4"},
	{"quote 5","author 5"},
}
func main() {
	rand.Seed(time.Now().Unix())

	api := echo.New()
	
	api.GET("/",getQuotes)
	api.GET("/quotes/random",getRandomQuote)

	port :=os.Getenv("PORT")
	if port == " " {
		port = "32444"
	}

	api.Start(":" + port)
}

func getQuotes(c echo.Context) error {
	return c.JSON(http.StatusOK,quotes)
}

func getRandomQuote(c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK,quotes[index])
}