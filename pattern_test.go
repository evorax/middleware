package middleware

import (
	"fmt"
	"log"
	"testing"
)

func TestPattern(t *testing.T) {
	param, err := MatchPath("/:id", "/test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(param)
}
