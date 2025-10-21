package mod

import (
	"fmt"

	"go.bergrend.dev/mod/mod1"
	"go.bergrend.dev/mod/mod2"
)

func Mod() {
	fmt.Println("mod.Mod()")
}

func All() {
	fmt.Println("mod.All()")
	mod1.Hello()
	mod2.Hello()
}
