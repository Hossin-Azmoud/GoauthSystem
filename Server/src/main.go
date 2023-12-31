package main

import (
    "fmt"
	"log"
	"github.com/Hossin-Azmoud/login_system/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


// TODO: LOAD FROM .env file.
const PORT = ":9090";
const HTML_CONTENT_DIR = "*.html";
const DB_PATH = ...;

func RequestCancelRecover() gin.HandlerFunc {
	
  return func(c *gin.Context) {
		defer func() {

			if err := recover(); err != nil {
				fmt.Println("The Request was cancelled because an unexpected error interupted.")
				fmt.Println("err:\n")
				log.Fatal(err);
				
				c.Request.Context().Done()
			}

		}()
		
		c.Next()
	}	
}


func run() {
	
	gin.SetMode(gin.ReleaseMode)
	/*	
	router.Static("/static", "./public/static")
	router.Static("/img", "./public/img")
	router.Static("/Global.css", "./public/Global.css")
	*/
	
	router := gin.Default()
	
	router.LoadHTMLGlob(HTML_CONTENT_DIR)
	router.Use(cors.Default())
	router.Use(gin.Logger(), RequestCancelRecover())
	
	router.GET("/", routes.Home)
	router.NoRoute(routes.Home)
	
	fmt.Println("Server is running on port: ", PORT);
	router.Run(PORT)
}

func main() {
  run();
}
