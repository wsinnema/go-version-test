package lib

import (
	"fmt"
	"github.com/wsinnema/go-version-test/subpackage"
)

func PackageVersion() {
	subpackage.SubPackageVersion()
	fmt.Println("In between versions")
}
