package main
import "C"
import "github.com/alist-org/alist/v3/cmd"

func main(){}
//export alist
func alist() {
	cmd.OutAlistInit()
}
