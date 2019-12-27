package flovm

import (
	"flo/vm"
	"fmt"
)

func main1() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	o := vm.Object{
		Constants: []vm.FloObject{
			vm.FloInteger(3),
			vm.FloInteger(2),
		},
		ConstantCount: 2,
		Names:         []string{},
		NameCount:     0,
		Instructions: []byte{
			vm.LOAD_CONST, 0,
			vm.LOAD_CONST, 1,
			vm.BINARY_ADD,
		},
		ObjectCode: []byte{},
	}
	flovm := vm.VM{}
	// o.Save("test.flobject")
	flovm.Init()
	flovm.Run(o)

}
