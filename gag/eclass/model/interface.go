package model

type EclassService interface {
	Login(body *LoginBody) (string, error)
}
