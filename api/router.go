package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var apiUrl = "https://api.telegram.org"

func init() {
	router = gin.Default()
	router.Any("/*path", func(context *gin.Context) {
		uri := context.Param("path")
		query := context.Request.URL.RawQuery
		url := apiUrl + uri + "?" + query
		fmt.Println(url)
		req, err := http.NewRequestWithContext(context, context.Request.Method, url, context.Request.Body)
		if err != nil {
			fmt.Println(err)
			context.String(http.StatusBadRequest, err.Error())
			return
		}
		req.Header = context.Request.Header
		req.PostForm = context.Request.PostForm
		req.Form = context.Request.Form
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			context.String(http.StatusBadRequest, err.Error())
			return
		}
		context.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, nil)
	})
}

func Listen(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
