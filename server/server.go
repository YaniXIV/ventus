package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	fmt.Println("Server starting!")
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	router.Run()
}

func initGSMerkleTree(c *gin.Context) {
	// TODO: Implement merkle tree initialization endpoint
}

/*
Construction algorithm notes:
- (1,2), (3,4), (5, 6) - pair leaves
- layer one is pre hashed
- find leaf, if even, hash to the left, else hash right. make sure not size 0 or 1.
- we are going to handle duplicating and stuff in the construction, don't worry about that
  in proof generation.
*/
