package utils

import (
	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(userId int, username, token, clientId, uuid string) (string, error) {

	claims := Claims{
		UserId: userId,
		//Username: username,
		//Password: password,
		AccessToken: token,
		ClientId:    clientId,
		//Uuid:     uuid,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: 0,
			// 指定token发行人
			Issuer: "baozier",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
