package main

import (
	"os"
)

func getResolverPath() string {
	return os.Getenv("RESOLVER_PATH")
}
