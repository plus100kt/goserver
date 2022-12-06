package model

type (
	Device struct {
		// 디바이스 ID
		UUID string `db:"uuid" json:"uuid"`

		// RSA Key
		RsaPrivateKey string `db:"rsa_private_key" json:"rsa_private_key"`
		RsaPublicKey  string `db:"rsa_public_key" json:"rsa_public_key"`
	}
)
