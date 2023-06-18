package database;

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
    "github.com/Hossin-Azmoud/login_system/models"
    "github.com/Hossin-Azmoud/login_system/security"
)

/*
DATABASE SCHEMA:
	CREATE USERS(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		EMAIL TEXT,
		USERNAME TEXT,
		PASSWORD_H TEXT,
	);
*/

var db *sql.DB
func InitializeDb(path string) error {
	
	var err error

	db, err = sql.Open("sqlite3", path); if err != nil {
		return err 
	}

	return nil;
}

func CheckUserExistence(Email string) bool {
	row, err := DATABASE.Query("SELECT ID FROM USERS WHERE EMAIL=? ORDER BY ID DESC", Email)
	defer row.Close()

	var u []int;
	
	if err != nil {
		fmt.Println(err)
		return false
	}

	var id int;

	for row.Next() {	
		row.Scan(&id)
		u = append(u, id)
	}

	return (len(u) >= 1);
}

func AddUser(user *models.User) (error) {
	
	if !CheckUserExistence(user.Email) {
		hash := security.Sha256_(user.Password)
		stmt, _ := DATABASE.Prepare("INSERT INTO USERS(EMAIL, USERNAME, PASSWORD_H) VALUES(?, ?, ?)")
		_, err := stmt.Exec(user.Email, user.UserName, hash)
		
		if err != nil {
			return err;
		}	
		
		Jwt, err := security.StoreUserInJWT(user.Email, user.Password)
		if err != nil {
			fmt.Println(err)
			return err;
		}
		
		return nil;
	}
	
	return errors.New("This user already exists..")
}

func AuthUser(user models.User) {
	
}

func GetUserInfo(user *models.User) bool {
	row, err := DATABASE.Query("SELECT ID, USERNAME FROM USERS WHERE EMAIL=? ORDER BY ID DESC", user.Email)

	defer row.Close()
	
	if err != nil {
		fmt.Println(err)
		return false
	}

	for row.Next() {	
		row.Scan(&user.id, &user.UserName)
	}

	return true
}

func AuthenticateUserJWT(user_jwt string) (models.User, error) {
    
	user, Ok := security.GetTokenFromJwt(user_jwt);
    
	if !Ok {
		return errors.New("Invalid JWT.")
	}
	
	ok := GetUserInfo(&user);
	
	if ok {	
		return user, nil;
	}

	return user, errors.New("User does not exist!")
}

func AuthenticateUserByEmailAndPwd(Pwd string, Email string) (models.User, error) {
	
	var user models.User

	if !CheckUserExistence(Email) {

		var user models.User
		row, err := DATABASE.Query("SELECT ID, EMAIL, USERNAME, PASSWORD_H FROM USERS WHERE EMAIL=?", Email)
		
		defer row.Close()

		if err != nil {
			fmt.Println(err)
			return models.User{}, errors.New("Could not get user from db. 82")
		}

		var pwdHash string

		for row.Next() {
			row.Scan(&pwdHash)
		}
		
		if security.Sha256_(Pwd) == pwdHash {
			row, err := DATABASE.Query("SELECT ID, EMAIL, USERNAME FROM USERS WHERE EMAIL=? ORDER BY ID DESC", Email)

			defer row.Close()
			if err != nil {
				return user, models.MakeServerResult(false, "Could not get user from db. 97")
			}

			for row.Next() {
				row.Scan(&user.Id_, &user.Email, &user.UserName, &user.Token, &user.Img, &user.Bg,  &user.Bio, &user.Address)
				user.Img = CheckCdnLink(user.Img);
				user.Bg = CheckCdnLink(user.Bg);
			}

			JWT, err := crypto.StoreTokenInJWT(user.Token)

			if err == nil {
				user.Token = JWT
				return user, models.MakeServerResult(true, "User created! you can login now..")
			}


			return EmptyUser, models.MakeServerResult(false, "Server had a problem encoding the token..")
		}
		
		return user, models.MakeServerResult(false, "incorrect password. try again")
	}

	return EmptyUser, errors.New("Could not find the wanted user.")
}


