package configurations

import (
	"github.com/gbrlsnchs/jwt/v3"
	log "github.com/sirupsen/logrus"
)

type CustomPayload struct {
	jwt.Payload
}

func GetToken() string {

	key := "sirion"
	issuer := "s3crtF9Z8K19ftE0Ces"

	var hs = jwt.NewHS256([]byte(key))
	pl := CustomPayload{
		Payload: jwt.Payload{
			Issuer: issuer,
		},
	}
	token, err := jwt.Sign(pl, hs)
	if err != nil {
		log.Error("Error in creating JWT token: ", err)
	}

	return string(token)

}
