package service

import (
	"net/http"

	"github.com/plus100kt/goserver/gag/eclass/model"
)

type eclassService struct {
	cookies []*http.Cookie
}

type EclasssConfig struct {
	cookies []*http.Cookie
}

func NewEclassService(c *EclasssConfig) model.EclassService {
	return &eclassService{
		cookies: c.cookies,
	}
}
