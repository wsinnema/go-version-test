package lib

import (
	"fmt"
	"github.com/wsinnema/go-version-test/subpackage"
)

func PackageVersion() {
	subpackage.SubPackageVersion()
	fmt.Println("Better than 4, but worse than 5")
}
