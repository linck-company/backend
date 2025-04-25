package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GetNewULID() string {
	source := rand.NewSource(time.Now().UnixNano())
	entropy := rand.New(source)
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}
