package controllers

import (
	"crud/models"
	"fmt"
)

func DeleteUser(users []models.User) []models.User {
	var userId int

	ListUsers(users)

	fmt.Print("# Digite o id do usuário: ")
	fmt.Scan(&userId)

	if userId > len(users)-1 {
		fmt.Println("# Usuário não encontrado!!!")
		return nil
	}

	fmt.Println("# Usuário foi deletado!!!")
	return append(users[:userId], users[userId+1:]...)
}
