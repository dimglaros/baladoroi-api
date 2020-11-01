package server

import (
	"fmt"
	"github.com/dimglaros/baladoroi-api/pkg/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func Run() {

	fmt.Println("Application booting...")

	var s Server

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	dsn := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=True&loc=Local"
	err = s.Initialize(dsn)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8000", s.Router))
}

func (s *Server) Initialize(dsn string) error {
	var err error
	s.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = s.DB.Debug().AutoMigrate(&models.User{}, &models.Field{}, &models.Game{}, &models.Participation{})
	if err != nil {
		return err
	}

	s.Router = mux.NewRouter()

	s.InitializeRoutes()

	return nil
}
