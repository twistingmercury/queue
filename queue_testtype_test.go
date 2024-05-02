package queue_test

import (
	"fmt"
	"math/rand"
	"time"
)

type testType struct {
	id         int64
	FirstName  string
	LastName   string
	Email      string
	Birthday   string
	Properties []any
}

func newTestType() testType {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	id := rng.Int63()
	firstName := generateRandomString(rng, 8)
	lastName := generateRandomString(rng, 10)
	email := fmt.Sprintf("%s.%s@example.com", firstName, lastName)
	birthday := fmt.Sprintf("%04d-%02d-%02d", rng.Intn(100)+1900, rng.Intn(12)+1, rng.Intn(28)+1)

	numProperties := rng.Intn(10) + 1
	properties := make([]any, numProperties)
	for i := 0; i < numProperties; i++ {
		propertyType := rng.Intn(4)
		switch propertyType {
		case 0:
			properties[i] = rng.Int()
		case 1:
			properties[i] = generateRandomString(rng, 12)
		case 2:
			properties[i] = rng.Float64()
		case 3:
			properties[i] = rng.Int63n(2) == 0
		}
	}

	return testType{
		id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Birthday:   birthday,
		Properties: properties,
	}
}

func generateRandomString(rng *rand.Rand, length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}
