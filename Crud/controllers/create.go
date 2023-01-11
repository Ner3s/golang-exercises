package controllers

import (
	"crud/models"
	"fmt"
)

func CreateUser() models.User {
	var user models.User

	fmt.Print("Insira o nome do usuário: ")
	fmt.Scan(&user.Name)

	fmt.Print("Insira o sobrenome do usuário: ")
	fmt.Scan(&user.LastName)

	fmt.Print("Insira o e-mail do usuário: ")
	fmt.Scan(&user.Email)

	fmt.Print("Insira o salário do usuário: ")
	fmt.Scan(&user.Salary)

	return user
}
