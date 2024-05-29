package main

import (
	"log"
	"net/http"

	"github.com/Jhonny1124/Renta_Carros/Backend/controllers"
	"github.com/Jhonny1124/Renta_Carros/Backend/handlers"
	"github.com/Jhonny1124/Renta_Carros/Backend/models"
	"github.com/Jhonny1124/Renta_Carros/Backend/repository"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func main() {

	conn, err := ConectarDB("joda", "postgres")
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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /cars", handler.ListarCarros())
	mux.HandleFunc("GET /cars/marca/{marca}", handler.ListarCarrosMarca())
	mux.HandleFunc("GET /cars/tipo/{tipo}", handler.ListarCarrosTipo())
	mux.HandleFunc("GET /cars/transmision/{transmision}", handler.ListarCarrosTransmision())
	mux.HandleFunc("GET /cars/combustible/{combustible}", handler.ListarCarrosCombustible())
	mux.HandleFunc("POST /cars", handler.CrearCarros())
	mux.HandleFunc("GET /cars/id/{id}", handler.TraerCarro())
	mux.HandleFunc("PATCH /cars/{id}", handler.ActualizarCarro())
	mux.HandleFunc("DELETE /cars/{id}", handler.EliminarCarro())

	mux.HandleFunc("GET /users", handler2.ListarUsuarios())
	mux.HandleFunc("POST /users", handler2.CrearUsuario())
	mux.HandleFunc("GET /users/{id}", handler2.TraerUsuario())
	mux.HandleFunc("PATCH /users/{id}", handler2.ActualizarUsuario())
	mux.HandleFunc("DELETE /users/{id}", handler2.EliminarUsuario())

	log.Fatal(http.ListenAndServe(":8080", mux))
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
