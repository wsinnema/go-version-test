package main

import (
	"fmt"
	"github.com/wsinnema/go-version-test/subpackage"
)

func main() {
	subpackage.SubPackageVersion()
	fmt.Println("In between versions")
}
