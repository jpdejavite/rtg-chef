package log

import (
	"math/rand"
	"time"
)

// CoiConfig config to generate a custom coi
type CoiConfig struct {
	Seed    int64
	Charset string
	Length  int
}

const (
	defaultCharset = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	defaultLength  = 15
)

// GenerateCoi generate a new random coi (CoifConfig is not required)
func GenerateCoi(coiConfig *CoiConfig) string {
	config := CoiConfig{
		Charset: defaultCharset,
		Length:  defaultLength,
		Seed:    time.Now().UnixNano(),
	}
	if coiConfig != nil {
		config = *coiConfig
	}

	return stringWithCharset(config.Length, config.Charset, config.Seed)
}

func stringWithCharset(length int, charset string, seed int64) string {
	seededRand := rand.New(rand.NewSource(seed))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
