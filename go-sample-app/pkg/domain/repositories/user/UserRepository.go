package repositories

import (
	"log"

	"go-sample-app/pkg/domain/models"
	"go-sample-app/pkg/infrastructure"
)

// 本番用のリポジトリ
type UserRepository struct {
}

func (rep *UserRepository) GetAllUsers() []models.User {
	// まだコネクトを都度行う
	db, err := infrastructure.ConnectDB()
	rows, err := db.Query("select id, name from users")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func (rep *UserRepository) FindUserById(id string) models.User {
	var user models.User

	// まだコネクトを都度行う
	db, err := infrastructure.ConnectDB()
	rows, err := db.Query("select id, name from users where ID = ?", id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return user
}
