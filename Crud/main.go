package main

import (
	"bufio"
	"crud/controllers"
	"crud/models"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var operator int = 1

	var users []models.User

	for operator != 0 {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println("################################")
		fmt.Println("#      GESTÃO DE USUÁRIOS      #")
		fmt.Println("################################")

		fmt.Println("Operadores: ")
		fmt.Println("1 - Adicionar novo usuário")
		fmt.Println("2 - Editar usuário")
		fmt.Println("3 - Listar usuários")
		fmt.Println("4 - Deletar usuário")
		fmt.Println("0 - SAIR")
		fmt.Print("Opção: ")
		fmt.Scan(&operator)

		switch operator {
		case 1:
			users = append(users, controllers.CreateUser())
		case 2:
			controllers.EditUser(users)
		case 3:
			controllers.ListUsers(users)
		case 4:
			{
				resultDelete := controllers.DeleteUser(users)
				if resultDelete != nil {
					users = resultDelete
				}
				break
			}
		case 0:
			fmt.Println("Obrigado!!!")
		default:
			fmt.Println("Opção selecionada é inválida!!! :( ")
		}

		fmt.Println("\n# pressione Enter para prosseguir!")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

}
