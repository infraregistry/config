package config

import (
	"crypto/ecdsa"
	"os"

	"github.com/nvr-ai/go-util/security"
)

type C struct {
	JwtPrivateKey *ecdsa.PrivateKey
	JwtPublicKey  *ecdsa.PublicKey
}

var Config *C = &C{}

func Init() {
	if os.Getenv("JWT_PRIVATE_KEY") != "" && os.Getenv("JWT_PUBLIC_KEY") != "" {
		jwtPrivateKey, err := security.DecodePrivateKey(os.Getenv("JWT_PRIVATE_KEY"))
		if err != nil {
			panic(err)
		}

		jwtPublicKey, err := security.DecodePublicKey(os.Getenv("JWT_PUBLIC_KEY"))
		if err != nil {
			panic(err)
		}

		Config.JwtPrivateKey = jwtPrivateKey
		Config.JwtPublicKey = jwtPublicKey
	}
}
