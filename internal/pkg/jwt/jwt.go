package jwt

import (
	"log"
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/joho/godotenv"
)

var SecretKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Unable to load .env file")
	}
	SecretKey = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(username string) (string,error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username 
	claims["exp"] = time.Now().Add(time.Hour*24).Unix()
	tokenString,err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating Key") 
		return "",err
	}
	return tokenString,nil 
}

func ParseToken(tokenStr string) (string,error){
	token,err := jwt.Parse(tokenStr,func(token *jwt.Token) (any,error){
		return SecretKey,nil 
	})
	if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		username := claims["username"].(string)
		return username,nil 
	}else{
		return "",err
	}
}