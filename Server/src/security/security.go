package crypto;

import (
    "crypto/sha256"
    b64 "encoding/base64"
    "fmt"
    "encoding/hex"
    "github.com/golang-jwt/jwt"
    "math/rand"
    "time"
    "strconv"
    "github.com/Hossin-Azmoud/login_system/loaders"
    "github.com/Hossin-Azmoud/login_system/models"
)

var SERVER_KEY []byte = []byte(loaders.GetEnv("JWT_KEY"))

/* DONE */

func Sha256_(s string) string {
    /* gets a string and returned a hash using sh256 */
    hash_ := sha256.New()
    hash_.Write([]byte(s))
    return hex.EncodeToString(hash_.Sum(nil))
}

// Store User info in the JWT. 
func StoreUserInJWT(Email string, Password string) (string, error) {
    
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": Email,
		"Passord": Password,
    })

    tokenString, err := token.SignedString(SERVER_KEY)
    return tokenString, err

}

//Verify JWT
func VerifyJWT(TokenStr string) (string, bool) {
	
	token, err := jwt.Parse(TokenStr, func(token *jwt.Token) (interface{}, error) { 
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return JWT_SECRET_KEY, nil
    })
	
	return token, (err != nil)
}

func GetUserFromJwt(TokenStr string) (models.User, bool) { 
    
	var User_Object models.User 
	token, ok := VerifyJWT(TokenStr)
	
	if !ok { 
        return User_Object, false
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {		
		// TODO: Define the data to be stored/returned from the jwt
		
		Email, _ := claims["Email"].(string)
		Password, _ := claims["Email"].(string)
		User_Object.Email = Email
		User_Object.Password = Password
		User_Object.Jwt = TokenStr		
		return User_Object, true
    }
    
	return User_Object, false
}
