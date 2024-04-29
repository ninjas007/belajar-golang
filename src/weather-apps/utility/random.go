package utility

import (
	"math/rand"
	"time"
)

// fungsi untuk generate random berdasarkan jarak angka minimal dan maksimal// fungsi untuk generate random untuk golang versi 1.21
func GenerateRandomNumber(min int, max int) (number int) {

	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	// proses melakukan random data berdasarkan angka minimum dan maksimum.
	return generator.Intn(max-min) + min
}
