package tablet

//IsTabletUserAgent - Returns true if the UserAgent is a tablet browser.
func IsTabletUserAgent(ua string) bool {
	return tabletAgent(ua)
}
