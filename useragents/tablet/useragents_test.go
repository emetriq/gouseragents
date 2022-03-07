package tablet

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTabletUserAgent(t *testing.T) {
	assert.True(t, IsTabletUserAgent("samsung SM-T585 Android 8.1.0 tablet"))
	assert.False(t, IsTabletUserAgent("5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)  5.122092"))
}

func TestIsTabletUserAgentJson(t *testing.T) {
	var userAgents = []string{}
	byteValue, _ := ioutil.ReadFile("../gen/ua_tablet.json")
	json.Unmarshal(byteValue, &userAgents)
	for _, ua := range userAgents {
		assert.True(t, IsTabletUserAgent(ua))
	}
}

func TestIsNotTabletUserAgentJson(t *testing.T) {
	files := []string{"../gen/ua_ctv.json", "../gen/ua_desktop.json", "../gen/ua_phone.json"}
	for _, file := range files {
		var userAgents = []string{}
		byteValue, _ := ioutil.ReadFile(file)
		json.Unmarshal(byteValue, &userAgents)
		for _, ua := range userAgents {
			assert.Falsef(t, IsTabletUserAgent(ua), "ua %s from file %s should not match as desktop user agent", ua, file)
		}
	}
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
