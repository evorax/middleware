package middleware

import (
	"fmt"
	"log"
	"testing"

	"github.com/evorax/middleware/internal"
)

func TestPattern(t *testing.T) {
	param, err := internal.Match("/:id", "/test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(param)
}
