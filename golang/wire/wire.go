package wire

// My External Storage "google/wireを使ったDIとDI関数のシグネチャについて #go"
// https://budougumi0617.github.io/2018/12/14/how-to-use-google-wire/

import (
	"context"
	"github.com/pkg/errors"
)

type Foo struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 42}
}

type Bar struct {
	X int
}

func ProvideBar(foo Foo) Bar {
	return Bar{X: -foo.X}
}

type Baz struct {
	X int
}

func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}

	return Baz{X: bar.X}, nil
}

