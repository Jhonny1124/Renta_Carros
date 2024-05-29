package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Jhonny1124/Renta_Carros/Backend/controllers"
)

type HandlerCarros struct {
	controller *controllers.Controller
}

type HandlerUsuarios struct {
	controller *controllers.ControllerU
}

func NewHandlerCarros(controller *controllers.Controller) (*HandlerCarros, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &HandlerCarros{
		controller: controller,
	}, nil
}

func NewHandlerUsuarios(controller *controllers.ControllerU) (*HandlerUsuarios, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &HandlerUsuarios{
		controller: controller,
	}, nil
}

func (hc *HandlerCarros) ListarCarros() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carros, err := hc.controller.ListarCarros()
		if err != nil {
			log.Printf("fallo al leer carros, con erro: %s", err.Error())
			http.Error(w, "fallo al leer carros", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(carros)
	})
}
func (hc *HandlerCarros) ListarCarrosMarca() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carros, err := hc.controller.ListarCarrosMarca()
		if err != nil {
			log.Printf("fallo al leer carros por marca, con erro: %s", err.Error())
			http.Error(w, "fallo al leer carros por marca", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(carros)
	})
}
func (hc *HandlerCarros) ListarCarrosTipo() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carros, err := hc.controller.ListarCarrosTipo()
		if err != nil {
			log.Printf("fallo al leer carros por tipo, con erro: %s", err.Error())
			http.Error(w, "fallo al leer carros por tipo", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(carros)
	})
}

func (hc *HandlerCarros) ListarCarrosTransmision() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carros, err := hc.controller.ListarCarrosTransmision()
		if err != nil {
			log.Printf("fallo al leer carros por transmision, con erro: %s", err.Error())
			http.Error(w, "fallo al leer carros por transmision", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(carros)
	})
}

func (hc *HandlerCarros) ListarCarrosCombustible() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carros, err := hc.controller.ListarCarrosCombustible()
		if err != nil {
			log.Printf("fallo al leer carros por combustible, con erro: %s", err.Error())
			http.Error(w, "fallo al leer carros por combustible", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(carros)
	})
}

func (hc *HandlerUsuarios) ListarUsuarios() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usuarios, err := hc.controller.ListarUsuarios()
		if err != nil {
			log.Printf("fallo al leer usuarios, con erro: %s", err.Error())
			http.Error(w, "fallo al leer usuarios", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(usuarios)
	})
}

func (hc *HandlerCarros) CrearCarros() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("fallo al crear nuevo carro, con erro: %s", err.Error())
			http.Error(w, "fallo al crear nuevo carro", http.StatusBadRequest)
			return
		}
		nuevoId, err := hc.controller.CrearCarro(body)
		if err != nil {
			log.Printf("fallo al crear nuevo carro, con erro: %s", err.Error())
			http.Error(w, "fallo al crear nuevo carro", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, fmt.Sprintf("id del nuevo carro: %d", nuevoId))
	})
}
func (hc *HandlerUsuarios) CrearUsuario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("fallo al crear nuevo usuario, con erro: %s", err.Error())
			http.Error(w, "fallo al crear nuevo usuario", http.StatusBadRequest)
			return
		}
		nuevoId, err := hc.controller.CrearUsuario(body)
		if err != nil {
			log.Printf("fallo al crear nuevo usuario, con erro: %s", err.Error())
			http.Error(w, "fallo al crear nuevo usuario", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, fmt.Sprintf("id del nuevo usuario: %d", nuevoId))
	})
}

func (hc *HandlerCarros) TraerCarro() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}

		carro, err := hc.controller.TraerCarro(id)
		if err != nil {
			log.Printf("fallo a leer un carro, con error: %s", err.Error())
			http.Error(w, "fallo a leer un carro", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(carro)
	})
}

func (hc *HandlerUsuarios) TraerUsuario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}

		usuario, err := hc.controller.TraerUsuario(id)
		if err != nil {
			log.Printf("fallo a leer un usuario, con error: %s", err.Error())
			http.Error(w, "fallo a leer un usuario", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(usuario)
	})
}

func (hc *HandlerCarros) ActualizarCarro() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("fallo al actualizar un carro, con erro: %s", err.Error())
			http.Error(w, "fallo al actualizar un carro", http.StatusInternalServerError)
		}
		err = hc.controller.ActualizarCarro(body, id)
		if err != nil {
			log.Printf("fallo al actualizar un carro, con erro: %s", err.Error())
			http.Error(w, "fallo al actualizar un carro", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
}

func (hc *HandlerUsuarios) ActualizarUsuario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("fallo al actualizar un usuario, con erro: %s", err.Error())
			http.Error(w, "fallo al actualizar un usuario", http.StatusInternalServerError)
		}
		err = hc.controller.ActualizarUsuario(body, id)
		if err != nil {
			log.Printf("fallo al actualizar un usuario, con erro: %s", err.Error())
			http.Error(w, "fallo al actualizar un usuario", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
}

func (hc *HandlerCarros) EliminarCarro() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		err = hc.controller.EliminarCarro(id)
		if err != nil {
			log.Printf("fallo al Eliminar un carro, con erro: %s", err.Error())
			http.Error(w, "fallo al Eliminar nuevo carro", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
}

func (hc *HandlerUsuarios) EliminarUsuario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "No se encontro un id valido", http.StatusBadRequest)
			return
		}
		err = hc.controller.EliminarUsuario(id)
		if err != nil {
			log.Printf("fallo al Eliminar un usuario, con erro: %s", err.Error())
			http.Error(w, "fallo al Eliminar nuevo usuario", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
}
