package phone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPhoneUserAgent(t *testing.T) {
	assert.False(t, IsPhoneUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0"))
	assert.True(t, IsPhoneUserAgent("5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)  5.122092"))
}

func BenchmarkUserAgentIn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPhoneUserAgent("5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)  5.122092")
	}
}

func BenchmarkUserAgentNotIn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPhoneUserAgent("Mipad")
	}
}
