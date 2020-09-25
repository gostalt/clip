package password

import (
	"math/rand"
	"time"
)

var (
	digits   string = "0123456789"
	specials string = "~=+%^*/()[]{}/!@#$?|"
	upper    string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower    string = "abcdefghijklmnopqrstuvwxyz"
)

type GenerateOptions struct {
	Digits   bool
	Specials bool
	Upper    bool

	Length int
}

func Generate(opts GenerateOptions) string {
	rand.Seed(time.Now().UnixNano())

	// By default, add all lower characters
	enabled := lower

	// Make it impossible to generate an empty password
	if opts.Length == 0 {
		opts.Length = 1
	}

	if opts.Upper {
		enabled = enabled + upper
	}

	if opts.Digits {
		enabled = enabled + digits
	}

	if opts.Specials {
		enabled = enabled + specials
	}

	buf := make([]byte, opts.Length)
	for i := 0; i < opts.Length; i++ {
		buf[i] = enabled[rand.Intn(len(enabled))]
	}

	rand.Shuffle(len(buf), func(i int, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}
