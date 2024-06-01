package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jhonny1124/Renta_Carros/Backend/controllers"
	"github.com/Jhonny1124/Renta_Carros/Backend/handlers"
	"github.com/Jhonny1124/Renta_Carros/Backend/models"
	"github.com/Jhonny1124/Renta_Carros/Backend/repository"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func main() {
	conn, err := ConectarDB(fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")), "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
	}

	bd, err := repository.NewRepository[models.Carro](conn)
	if err != nil {
		log.Fatalln("error creando una instancia de repository", err.Error())
	}

	controller, err := controllers.Newcontroller(bd)
	if err != nil {
		log.Fatalln("error creando un nuevo controlador", err.Error())
	}

	handler, err := handlers.NewHandlerCarros(controller)
	if err != nil {
		log.Fatalln("error creando un nuevo handler", err.Error())
	}

	bdu, err := repository.NewRepository[models.Usuario](conn)
	if err != nil {
		log.Fatalln("error creando una instancia de repository", err.Error())
	}
	controller2, err := controllers.NewcontrollerU(bdu)
	if err != nil {
		log.Fatalln("error creando un nuevo controlador", err.Error())
	}
	handler2, err := handlers.NewHandlerUsuarios(controller2)
	if err != nil {
		log.Fatalln("error creando un nuevo handler", err.Error())
	}

	router := mux.NewRouter()

	router.Handle("/cars", http.HandlerFunc(handler.ListarCarros)).Methods(http.MethodGet)
	router.Handle("/cars/marca/{marca}", http.HandlerFunc(handler.ListarCarrosMarca)).Methods(http.MethodGet)
	router.Handle("/cars/tipo/{tipo}", http.HandlerFunc(handler.ListarCarrosTipo)).Methods(http.MethodGet)
	router.Handle("/cars/transmision/{transmision}", http.HandlerFunc(handler.ListarCarrosTransmision)).Methods(http.MethodGet)
	router.Handle("/cars/combustible/{combustible}}", http.HandlerFunc(handler.ListarCarrosCombustible)).Methods(http.MethodGet)
	router.Handle("/cars", http.HandlerFunc(handler.CrearCarros)).Methods(http.MethodPost)
	router.Handle("/cars/id/{id}", http.HandlerFunc(handler.TraerCarro)).Methods(http.MethodGet)
	router.Handle("/cars/{id}", http.HandlerFunc(handler.ActualizarCarro)).Methods(http.MethodPatch)
	router.Handle("/cars/{id}", http.HandlerFunc(handler.EliminarCarro)).Methods(http.MethodDelete)

	router.Handle("/users", http.HandlerFunc(handler2.ListarUsuarios)).Methods(http.MethodGet)
	router.Handle("/users", http.HandlerFunc(handler2.CrearUsuario)).Methods(http.MethodPost)
	router.Handle("/users/{id}", http.HandlerFunc(handler2.TraerUsuario)).Methods(http.MethodGet)
	router.Handle("/users/{id}", http.HandlerFunc(handler2.ActualizarUsuario)).Methods(http.MethodPatch)
	router.Handle("/users/{id}", http.HandlerFunc(handler2.EliminarUsuario)).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}
