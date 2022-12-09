package eclass

import (
	"net/http"

	"github.com/plus100kt/goserver/gag/eclass/model"
)

type Eclass struct {
	cookies []*http.Cookie
}

func NewEclass(e *Eclass) model.Eclass {
	return &Eclass{
		cookies: e.cookies,
	}
}
