package middleware

import (
	"fmt"
	"log"
	"testing"

	"github.com/evorax/middleware/internal"
)

func TestPattern(t *testing.T) {
	param, err := internal.MatchPath(`/foo/:id[^[a-zA-Z]+$]`, "/foo/test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(param)
}
