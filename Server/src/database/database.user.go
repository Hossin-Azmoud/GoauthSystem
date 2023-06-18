package database;

import (
	"fmt"
    "github.com/Hossin-Azmoud/login_system/models"
	"github.com/Hossin-Azmoud/login_system/security"
)

func AuthenticateUserJWT(UserJWT string) (bool, models.User) {
    Token, Ok := security.GetTokenFromJwt(UserJWT);
    
    if Ok {
     
        User_, err := GetUserByToken(Token)
		if err != nil {
            // a db error.
            return models.MakeGenericServerResponse(500, "Db Error. (line 108).")
        }
        
		return Ok, User_
	}
    
	return models.MakeGenericServerResponse(500, "server could not decode the token. (line 117)")
}
