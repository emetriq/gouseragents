package desktop

//IsDesktopUserAgent - Returns true if the UserAgent is a desktop browser.
func IsDesktopUserAgent(ua string) bool {
	return desktopAgent(ua)
}
