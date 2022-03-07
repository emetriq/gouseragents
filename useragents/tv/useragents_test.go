package tv

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTVUserAgent(t *testing.T) {
	assert.True(t, IsTVUserAgent("Dalvik/2.1.0 (Linux; U; Android 5.1.1; AFTT Build/LVY48F) CTV"))
	assert.False(t, IsTVUserAgent("5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)  5.122092"))
}

func TestIsTVUserAgentJson(t *testing.T) {
	var userAgents = []string{}
	byteValue, _ := ioutil.ReadFile("../gen/ua_ctv.json")
	json.Unmarshal(byteValue, &userAgents)
	for _, ua := range userAgents {
		assert.True(t, IsTVUserAgent(ua))
	}
}

func TestIsNotTVUserAgentJson(t *testing.T) {
	files := []string{"../gen/ua_tablet.json", "../gen/ua_desktop.json", "../gen/ua_phone.json"}
	for _, file := range files {
		var userAgents = []string{}
		byteValue, _ := ioutil.ReadFile(file)
		json.Unmarshal(byteValue, &userAgents)
		for _, ua := range userAgents {
			assert.Falsef(t, IsTVUserAgent(ua), "ua %s from file %s should not match as desktop user agent", ua, file)
		}
	}
}

func BenchmarkUserAgentIn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsTVUserAgent("Dalvik/2.1.0 (Linux; U; Android 5.1.1; AFTT Build/LVY48F) CTV")
	}
}

func BenchmarkUserAgentNotIn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsTVUserAgent("Mipad")
	}
}
