package eclass

import (
	"net/http"

	"github.com/plus100kt/goserver/gag/eclass/model"
)

type eclass struct {
	cookies []*http.Cookie
}

type EclasssConfig struct {
	cookies []*http.Cookie
}

func NewEclass(c *EclasssConfig) model.Eclass {
	return &eclass{
		cookies: c.cookies,
	}
}
