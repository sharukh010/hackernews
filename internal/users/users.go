package users

import (
	"database/sql"
	"log"

	database "github.com/sharukh010/hackernews/internal/pkg/db/migrations/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	if err != nil {
		log.Panic(err.Error())
	}
	hashedPassword,err := HashPassword(user.Password)
	if err != nil {
		log.Panic(err.Error())
	}
	_,err = stmt.Exec(user.Username,hashedPassword)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HashPassword(password string) (string,error){
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes),err
}

func CheckPasswordHash(password,hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil 
}
func GetUserIdByUsername(username string) (int,error){
	stmt,err := database.Db.Prepare("SELECT ID FROM Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err.Error())
	}
	row := stmt.QueryRow(username)
	var Id int 
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows{
			log.Print(err.Error())
		}
		return 0,err 
	}
	return Id,nil 
}