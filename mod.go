package mod

import (
	"fmt"
	"go.bergrend.dev/mod/mod1"
)

func Go() {
	fmt.Println("mod.Go()")
}

func CallMod1Go() {
	callFunc("mod1.Go", mod1.Go)
}

func callFunc(s string, fn func()) {
	fmt.Printf("\n\nmod.CallFunc(%s) start\n", s)
	fn()
	fmt.Printf("mod.CallFunc(%s) end\n\n", s)
}
