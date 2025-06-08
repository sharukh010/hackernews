package links

import (
	"log"

	database "github.com/sharukh010/hackernews/internal/pkg/db/migrations/mysql"
	"github.com/sharukh010/hackernews/internal/users"
)

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *users.User `json:"user"`
}

func (link Link) Save() int64 {
	stmt,err := database.Db.Prepare("INSERT INTO LINKS(Title,Address,UserID) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err.Error())
	}
	res,err := stmt.Exec(link.Title,link.Address,link.User.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	id,err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:",err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Link {
	stmt,err := database.Db.Prepare("SELECT L.id,L.title,L.address,L.UserID,U.Username from Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	rows,err := stmt.Query()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()
	var links []Link
	var username string 
	var id string 
	for rows.Next(){
		var link Link 
		err := rows.Scan(&link.ID,&link.Title,&link.Address,&id,&username)
		if err != nil {
			log.Fatal(err.Error())
		}
		link.User = &users.User{
			ID:id,
			Username: username,
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
	}
	return links
}