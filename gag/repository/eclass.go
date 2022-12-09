package repository

import (
	"context"

	eclassModel "github.com/plus100kt/goserver/gag/eclass/model"
	"github.com/plus100kt/goserver/gag/model"
	"github.com/plus100kt/goserver/gag/util"
)

type eclassRepository struct {
	Eclass eclassModel.Eclass
}

type EclassConfg struct {
	Eclass eclassModel.Eclass
}

func NewEclassRepository(c *EclassConfg) model.EclassRepository {
	return &eclassRepository{
		Eclass: c.Eclass,
	}
}

func (r eclassRepository) Login(ctx context.Context, key string, u *model.User) error {
	// RSA 복호화
	rh := util.RSAHelper{}
	rh.PrivateFromStringPEM(u.RsaPrivateKey)

	aesKey, err := rh.DecryptString(key)
	if err != nil {
		return err
	}

	iv := util.PKCS5Padding([]byte(aesKey[0:8]), 16)
	password := util.AESDecrypt([]byte(u.AesPassword), []byte(aesKey), iv)

	body := &eclassModel.LoginBody{
		Usr_id:  u.ID,
		Usr_pwd: string(password),
	}

	// 로그인
	err = r.Eclass.Login(ctx, body)
	if err != nil {
		return err
	}

	return err
}

func (r eclassRepository) GetUser(ctx context.Context, u *model.User) error {
	// 로그인
	student, err := r.Eclass.GetStudent(ctx)
	if err != nil {
		return err
	}

	u = &model.User{
		Name:     student.Name,
		Email:    student.Email,
		ImageURL: student.ImageUrl,
	}

	return err
}
