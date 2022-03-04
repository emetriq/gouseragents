package tv

//IsTVUserAgent - Returns true if the UserAgent is a TV browser.
func IsTVUserAgent(ua string) bool {
	return tvAgent(ua)
}
