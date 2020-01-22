package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moethu/codenutrition/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/thinkerou/favicon"
)

// operatingDir current binary dir
var operatingDir string

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	operatingDir = dir

	docs.SwaggerInfo.Title = "Code Nutrition Service API"
	docs.SwaggerInfo.Description = "Code Nutrition Service"
	docs.SwaggerInfo.Version = "0.1"

	flag.Parse()
	log.SetFlags(0)

	router := gin.Default()
	port := ":80"
	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	router.Use(favicon.New("./static/favicon.ico"))
	router.Static("/static/", "./static/")
	router.GET("/", getHome)
	router.GET("/imprint", getImprint)
	router.GET("/facts/:code", getFacts)
	router.GET("/badge/:code", getBadge)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Starting HTTP Server on Port", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}

// getHome godoc
// @Summary renders form
// @Description renders form to generate code string
// @Produce  html
// @Success 200 {html} Status "OK"
// @Router / [get]
func getHome(c *gin.Context) {
	t := template.Must(template.ParseFiles("templates/form.html"))
	t.Execute(c.Writer, "http://"+c.Request.Host)
}

// getImprint godoc
// @Summary renders imprint
// @Description renders imprint
// @Produce  html
// @Success 200 {html} Status "OK"
// @Router /imprint [get]
func getImprint(c *gin.Context) {
	t := template.Must(template.ParseFiles("templates/imprint.html"))
	t.Execute(c.Writer, "http://"+c.Request.Host)
}

// Data struct for facts route
type Data struct {
	Host string
	Code string
}

// getFacts godoc
// @Summary renders facts page
// @Description renders facts page using code parameter string
// @Produce  html
// @Param code path string true "Code String"
// @Success 200 {html} Status "OK"
// @Router /facts/{code} [get]
func getFacts(c *gin.Context) {
	d := Data{Host: "http://" + c.Request.Host, Code: escapedParam(c, "code")}
	t := template.Must(template.ParseFiles("templates/result.html"))
	t.Execute(c.Writer, d)
}

// getBadge godoc
// @Summary renders badge image
// @Description renders badge image as png from code parameter string
// @Produce  image/png
// @Param code path string true "Code String"
// @Success 200 {string} Status "OK"
// @Router /badge/{code} [get]
func getBadge(c *gin.Context) {
	codelabel := escapedParam(c, "code")
	segments := strings.Split(codelabel, "_")
	image := createBadge(segments)
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(image)))
	if _, err := c.Writer.Write(image); err != nil {
		log.Println(err)
	}
}

// escapedParam gets a url parameter with escaped content and a max. length of 40
func escapedParam(c *gin.Context, param string) string {
	value := template.HTMLEscapeString(c.Param(param))
	if len(value) > 40 {
		return ""
	}
	return value
}
