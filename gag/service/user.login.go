package service

import (
	"context"

	"github.com/plus100kt/goserver/gag/model"
	"github.com/plus100kt/goserver/gag/util"
)

func (s *userService) Login(ctx context.Context, key string, u *model.User) (*model.User, error) {
	device, err := s.DeviceRepository.FindByID(ctx, u.UUID)
	if err != nil {
		return u, err
	}
	rsaPrivateKey := device.RsaPrivateKey

	// ------- id가 암호화 되어 있지 않다면 불필요 코드 -----
	// RSA 복호화
	rh := util.RSAHelper{}
	rh.PrivateFromStringPEM(rsaPrivateKey)

	aesKey, err := rh.DecryptString(key)
	if err != nil {
		return u, err
	}

	iv := util.PKCS5Padding([]byte(aesKey[0:8]), 16)
	// AES 복호화
	id := util.AESDecrypt([]byte(u.ID), []byte(aesKey), iv)
	// ------- id가 암호화 되어 있지 않다면 불필요 코드 -----

	u = &model.User{
		ID:            string(id),
		RsaPrivateKey: rsaPrivateKey,
	}

	u, err = s.EclassRepository.Login(ctx, key, u)
	if err != nil {
		return u, err
	}

	// 로그인 성공시 DB 저장
	err = s.UserRepository.Create(ctx, u)
	if err != nil {
		return u, err
	}

	err = s.DeviceRepository.Delete(ctx, u.UUID)
	if err != nil {
		return u, err
	}

	return u, err
}
