package mod

import (
	"fmt"
	"testing"

	"go.bergrend.dev/mod/mod1"
	"go.bergrend.dev/mod/mod2"
	"go.bergrend.dev/mod/quote"
)

func TestName(t *testing.T) {
	fmt.Println("mod.main()")
	Mod()
	mod1.Hello()
	mod2.Hello()
	quote.Go()
}
