package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Date string `form:"date"`
}

func main() {
	router := gin.Default()
	router.GET("/epoch", dateQuery)
	router.Run()
}

func dateQuery(c *gin.Context) {
	var query Query
	c.ShouldBind(&query)

	cmd := exec.Command("/bin/sh", "do.sh", query.Date)
	cmd.Dir = "."
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	var unixtimestamp = strings.TrimSuffix(out.String(), "\n")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Error ", err)
		return
	}

	fmt.Println("Timestamp ", unixtimestamp)

	c.JSON(http.StatusOK, gin.H{"timestamp": unixtimestamp})
}
