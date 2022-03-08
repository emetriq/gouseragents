[![Coverage Status](https://coveralls.io/repos/github/emetriq/gouseragents/badge.svg?branch=main)](https://coveralls.io/github/emetriq/gouseragents?branch=main) [![RELEASE](https://github.com/emetriq/gouseragents/actions/workflows/release.yaml/badge.svg)](https://github.com/emetriq/gouseragents/actions/workflows/release.yaml) [![CI](https://github.com/emetriq/gouseragents/actions/workflows/ci.yaml/badge.svg)](https://github.com/emetriq/gouseragents/actions/workflows/ci.yaml)

# gouseragents

Accurate and fresh list of desktop, phone, tablet and tv user agents.

# install

`go get github.com/emetriq/gouseragents`

# usage

```golang
import (
   "fmt"
   "github.com/emetriq/gouseragents/desktop"
   "github.com/emetriq/gouseragents/tv"
   "github.com/emetriq/gouseragents/phone"
   "github.com/emetriq/gouseragents/tablet"
)

func main() {
   if desktop.IsDesktopUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0") {
      fmt.Println("Desktop user agent")
   }

   if tv.IsTVUserAgent("Dalvik/2.1.0 (Linux; U; Android 9; AFTSSS Build/PS7242) CTV") {
      fmt.Println("TV user agent")
   }

   if phone.IsPhoneUserAgent("samsung SM-A515F Android 11 phone") {
      fmt.Println("Phone user agent")
   }

   if tablet.IsTabletUserAgent("samsung SM-T585 Android 8.1.0 tablet") {
      fmt.Println("Tablet user agent")
   }
}
```

# update useragents

The code base are json files located in the "gen" directory. After a file change you have to start the code generator with `go generate gen/useragents.go`. Create a PR. Done.
