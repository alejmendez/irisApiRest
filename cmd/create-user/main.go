package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alejmendez/goApiRest/app/model"
	"github.com/alejmendez/goApiRest/app/repositories"
	"github.com/alejmendez/goApiRest/app/services"
	"github.com/alejmendez/goApiRest/app/utils"
	"github.com/alejmendez/goApiRest/core/config"
	"github.com/alejmendez/goApiRest/core/database"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Definir flags
	name := flag.String("n", "", "Nombre del usuario")
	email := flag.String("e", "", "Email del usuario")
	password := flag.String("p", "", "Contraseña del usuario")
	flag.Parse()

	// Validar que se proporcionaron todos los argumentos requeridos
	if *name == "" || *email == "" || *password == "" {
		fmt.Println("Uso: go run cmd/create-user/main.go -n \"nombre\" -e \"email\" -p \"contraseña\"")
		flag.PrintDefaults()
		os.Exit(1)
	}

	config.InitConfig()
	// Conectar a la base de datos usando la conexión existente
	db := database.ConnectDB()

	// Inicializar repositorio y servicio
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Crear hash de la contraseña
	hashedPassword := utils.GenerateHash(*password)

	// Crear usuario
	user := &model.User{
		Username: *name,
		Email:    *email,
		Password: hashedPassword,
	}

	createdUser, err := userService.Create(user)
	if err != nil {
		log.Fatal("Error creando el usuario:", err)
	}

	fmt.Printf("Usuario creado exitosamente:\nID: %s\nUsername: %s\nEmail: %s\n",
		createdUser.ID, createdUser.Username, createdUser.Email)
}
