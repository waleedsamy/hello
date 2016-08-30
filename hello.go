package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
	"github.com/golang/glog"
)

func main() {

	glog.Info("Hello, world.\n")

	fmt.Printf(stringutil.Reverse("Hello, world.\n"))
}
