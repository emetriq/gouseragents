# gouseragents

Accurate and fresh list of desktop, phone, tablet and tv user agents.

# install

`go get github.com/emetriq/gouseragents`

# usage

```golang
import (
   desktop  "github.com/emetriq/gouseragents/desktop"
	"fmt"
)

func main() {
	if desktop.IsDesktopUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0") {
		fmt.Println("Desktop user agent")
	}
}
```

# update useragents

The code base are json files located in the "gen" directory. After a file change you have to start the code generator with `go generate gen/useragents.go`. Create a PR. Done.