package controllers

import (
	"crud/models"
	"fmt"
)

func ListUsers(users []models.User) {

	fmt.Println("#################################################################################################")
	fmt.Println("# LISTAGEM DE USUÁRIOS                                                                          #")
	fmt.Println("#################################################################################################")
	fmt.Println("#\t* ID * \t* NOME * \t* SOBRENOME * \t* EMAIL * \t\t\t* SALÁRIO * \t#")

	for key, value := range users {
		fmt.Printf("#\t  %d   \t  %s   \t  %s   \t  %s   \t  %s   \t#\n", key, value.Name, value.LastName, value.Email, value.Salary)
	}

	fmt.Println("\n#################################################################################################")
}
