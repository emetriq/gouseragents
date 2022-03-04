package tablet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTabletUserAgent(t *testing.T) {
	assert.True(t, IsTabletUserAgent("samsung SM-T585 Android 8.1.0 tablet"))
	assert.False(t, IsTabletUserAgent("5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)  5.122092"))
}

func BenchmarkUserAgentIn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsTabletUserAgent("samsung SM-T585 Android 8.1.0 tablet")
	}
}

func BenchmarkUserAgentNotIn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsTabletUserAgent("Mipad")
	}
}
