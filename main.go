package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rugby-board/go-rugby-dict/dict"
)

func main() {
	r := gin.Default()
	d := dict.NewDict("dict.yaml")
	err := d.Load()
	if err != nil {
		fmt.Println("Load dict failed")
		return
	}
	r.GET("/translate", func(c *gin.Context) {
		entry := c.Query("entry")
		trans, err := d.Query(entry)
		msg := "success"
		if err != nil {
			trans = "error"
		}
		c.JSON(200, gin.H{
			"entry":   entry,
			"result":  trans,
			"message": msg,
		})
	})
	r.Run()
}
