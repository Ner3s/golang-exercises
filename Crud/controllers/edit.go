package controllers

import (
	"crud/models"
	"fmt"
)

func EditUser(users []models.User) {
	var userId int

	ListUsers(users)

	fmt.Print("# Digite o id do usuário: ")
	fmt.Scan(&userId)

	if userId > len(users)-1 {
		fmt.Println("# Usuário não encontrado!!!")
		return
	}

	users[userId] = CreateUser()
}
