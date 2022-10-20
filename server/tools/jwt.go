package tools

import (
	"fmt"
	"github.com/HolisChen/tiny-blog/model"
	"github.com/golang-jwt/jwt/v4"
)

var (
	hmacSampleSecret []byte = []byte("HelloWorld")
)

func GenerateJwt(author *model.Author) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"loginId": author.LoginId,
		"mail":    author.MailAddress,
	})
	signedString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}
	return signedString
}

func ParseJwt(jwtStr string) (loginId string, mail string, err error) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		return "", "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["loginId"].(string), claims["mail"].(string), nil
	} else {
		return "", "", err
	}
}
