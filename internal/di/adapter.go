package di

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/handler"
)

var AdapterSet = wire.NewSet(NewAdapter)

type Adapter struct {
	Handler handler.Handler
}

func NewAdapter(h handler.Handler) Adapter {
	return Adapter{
		Handler: h,
	}
}
