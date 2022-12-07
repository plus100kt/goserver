package repository

import (
	"context"

	eclassModel "github.com/plus100kt/goserver/gag/eclass/model"
	"github.com/plus100kt/goserver/gag/model"
	"github.com/plus100kt/goserver/gag/util"
)

type eclassRepository struct {
	EclassService eclassModel.EclassService
}

type EclassConfg struct {
	EclassService eclassModel.EclassService
}

func NewEclassRepository(c *EclassConfg) model.EclassRepository {
	return &eclassRepository{
		EclassService: c.EclassService,
	}
}

func (r eclassRepository) Login(ctx context.Context, key string, u *model.User) (*model.User, error) {
	// RSA 복호화
	rh := util.RSAHelper{}
	rh.PrivateFromStringPEM(u.RsaPrivateKey)

	aesKey, err := rh.DecryptString(key)
	if err != nil {
		return u, err
	}

	iv := util.PKCS5Padding([]byte(aesKey[0:8]), 16)
	password := util.AESDecrypt([]byte(u.AesPassword), []byte(aesKey), iv)

	var LoginBody *eclassModel.LoginBody

	LoginBody = &eclassModel.LoginBody{
		Usr_id:  u.ID,
		Usr_pwd: string(password),
	}

	// 로그인
	string, err := r.EclassService.Login(LoginBody)

	// parsing

	return u, err
}
