package phone

//IsPhoneUserAgent - Returns true if the UserAgent is a mobile browser.
func IsPhoneUserAgent(ua string) bool {
	return phoneAgent(ua)
}
