package model

import (
	"time"
)

type (
	User struct {
		// 보안 계정 정보
		ID            string `db:"id" json:"id"`
		UUID          string `db:"uuid" json:"uuid"`
		RsaPrivateKey string `db:"rsa_private_key" json:"rsa_private_key"`
		AesPassword   string `db:"aes_password" json:"aes_password"`

		// 개인 정보
		Name     string `db:"name" json:"name"`
		Email    string `db:"email" json:"email"`
		ImageURL string `db:"image_url" json:"image_url"`
		Cookie   string `db:"cookie" json:"cookie"`

		// timestamp
		CreatedAt time.Time `db:"created_at" json:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	}
)
