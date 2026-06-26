package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateBarcode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%013d", r.Int63n(9000000000000)+1000000000000)
}
