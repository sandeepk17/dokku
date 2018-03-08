package main

import (
	"flag"

	"github.com/dokku/dokku/plugins/common"
)

// destroys the repo property for a given app container
func main() {
	flag.Parse()
	appName := flag.Arg(0)

	common.PropertyDestroy("repo", appName)
}
