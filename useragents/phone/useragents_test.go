package phone

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPhoneUserAgent(t *testing.T) {
	assert.False(t, IsPhoneUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0"))
	assert.True(t, IsPhoneUserAgent("5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)  5.122092"))
}

func TestIsPhoneUserAgentJson(t *testing.T) {
	var userAgents = []string{}
	byteValue, _ := ioutil.ReadFile("../gen/ua_phone.json")
	json.Unmarshal(byteValue, &userAgents)
	for _, ua := range userAgents {
		assert.True(t, IsPhoneUserAgent(ua))
	}
}

func TestIsNotPhoneUserAgentJson(t *testing.T) {
	files := []string{"../gen/ua_ctv.json", "../gen/ua_desktop.json", "../gen/ua_tablet.json"}
	for _, file := range files {
		var userAgents = []string{}
		byteValue, _ := ioutil.ReadFile(file)
		json.Unmarshal(byteValue, &userAgents)
		for _, ua := range userAgents {
			assert.Falsef(t, IsPhoneUserAgent(ua), "ua %s from file %s should not match as desktop user agent", ua, file)
		}
	}
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
