package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Jhonny1124/Renta_Carros/Backend/models"
	"github.com/Jhonny1124/Renta_Carros/Backend/repository"
)

var (
	ListQueryC   = "SELECT * FROM Carro"
	ListQueryCm  = "SELECT * FROM Carro WHERE marca=$1"
	ListQueryCt  = "SELECT * FROM Carro WHERE tipo=$1"
	ListQueryCtr = "SELECT * FROM Carro WHERE transmision=$1"
	ListQueryCc  = "SELECT * FROM Carro WHERE combustible=$1"
	ReadQueryC   = "SELECT * FROM Carro WHERE id=$1"
	CreateQueryC = "INSERT INTO Carro (marca, referencia, modelo, carroceria, potencia, torque, transmision, motor, pasajeros, combustible, consumo, almacenamiento, descripcion, lujo, deportivo, imagen) VALUES (:marca, :referencia, :modelo, :carroceria, :potencia, :torque, :transmision, :motor, :pasajeros, :combustible, :consumo, :almacenamiento, :descripcion, :lujo, :deportivo, :imagen) RETURNING id"
	UpdateQueryC = "UPDATE Carro SET %s WHERE id=:id"
	DeleteQueryC = "DELETE FROM Carro WHERE id=$1"
)
var (
	ListQueryU   = "SELECT * FROM Usuario"
	ReadQueryU   = "SELECT * FROM Usuario WHERE id=$1"
	CreateQueryU = "INSERT INTO Usuario (nombre, contrasena, reservas) VALUES (:nombre, :contrasena, :reservas) RETURNING id"
	UpdateQueryU = "UPDATE Usuario SET %s WHERE id=:id"
	DeleteQueryU = "DELETE FROM Usuario WHERE id=$1"
)

type Controller struct {
	repo repository.Repository[models.Carro]
}

type ControllerU struct {
	repo repository.Repository[models.Usuario]
}

func Newcontroller(repo repository.Repository[models.Carro]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para un controlador es necesario un repositorio valido")
	}
	return &Controller{
		repo: repo,
	}, nil
}
func NewcontrollerU(repo repository.Repository[models.Usuario]) (*ControllerU, error) {
	if repo == nil {
		return nil, fmt.Errorf("para un controlador es necesario un repositorio valido")
	}
	return &ControllerU{
		repo: repo,
	}, nil
}

func (c *Controller) ListarCarros() ([]byte, error) {
	carros, _, err := c.repo.List(context.Background(), ListQueryC, 200, 0)
	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}

	jsonCarro, err := json.Marshal(carros)

	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}
	return jsonCarro, nil
}

func (c *Controller) ListarCarrosMarca() ([]byte, error) {
	carros, _, err := c.repo.List(context.Background(), ListQueryCm, 200, 0)
	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}

	jsonCarro, err := json.Marshal(carros)

	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}
	return jsonCarro, nil
}

func (c *Controller) ListarCarrosTipo() ([]byte, error) {
	carros, _, err := c.repo.List(context.Background(), ListQueryCt, 200, 0)
	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}

	jsonCarro, err := json.Marshal(carros)

	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}
	return jsonCarro, nil
}
func (c *Controller) ListarCarrosTransmision() ([]byte, error) {
	carros, _, err := c.repo.List(context.Background(), ListQueryCtr, 200, 0)
	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}

	jsonCarro, err := json.Marshal(carros)

	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}
	return jsonCarro, nil
}

func (c *Controller) ListarCarrosCombustible() ([]byte, error) {
	carros, _, err := c.repo.List(context.Background(), ListQueryCc, 200, 0)
	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}

	jsonCarro, err := json.Marshal(carros)

	if err != nil {
		log.Printf("fallo a leer carros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer carros, con error: %s", err.Error())
	}
	return jsonCarro, nil
}

func (c *ControllerU) ListarUsuarios() ([]byte, error) {
	usuarios, _, err := c.repo.List(context.Background(), ListQueryU, 200, 0)
	if err != nil {
		log.Printf("fallo a leer usuarios, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer usuarios, con error: %s", err.Error())
	}

	jsonUsuarios, err := json.Marshal(usuarios)

	if err != nil {
		log.Printf("fallo a leer usuarios, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer usuarios, con error: %s", err.Error())
	}
	return jsonUsuarios, nil
}

func (c *Controller) TraerCarro(id string) ([]byte, error) {
	carro, err := c.repo.Read(context.Background(), ReadQueryC, id)
	if err != nil {
		log.Printf("fallo a leer un carro, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer un carro, con error: %s", err.Error())
	}
	Carrojson, err := json.Marshal(carro)
	if err != nil {
		log.Printf("fallo a leer un carro, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer un carro, con error: %s", err.Error())
	}
	return Carrojson, nil
}

func (c *ControllerU) TraerUsuario(id string) ([]byte, error) {
	usuario, err := c.repo.Read(context.Background(), ReadQueryU, id)
	if err != nil {
		log.Printf("fallo a leer un usuario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer un usuario, con error: %s", err.Error())
	}
	Usuariojson, err := json.Marshal(usuario)
	if err != nil {
		log.Printf("fallo a leer un usuario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo a leer un usuario, con error: %s", err.Error())
	}
	return Usuariojson, nil
}

func (c *Controller) CrearCarro(body []byte) (int64, error) {
	NuevoCarro := &models.Carro{}
	err := json.Unmarshal(body, NuevoCarro)
	if err != nil {
		log.Printf("fallo a crear un carro, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo a crear un carro, con error: %s", err.Error())
	}
	valores_columnas := map[string]any{
		"id":             NuevoCarro.Id,
		"marca":          NuevoCarro.Marca,
		"referencia":     NuevoCarro.Referencia,
		"modelo":         NuevoCarro.Modelo,
		"tipo":           NuevoCarro.Tipo,
		"potencia":       NuevoCarro.Potencia,
		"torque":         NuevoCarro.Torque,
		"transmision":    NuevoCarro.Transmision,
		"motor":          NuevoCarro.Motor,
		"pasajeros":      NuevoCarro.Pasajeros,
		"combustible":    NuevoCarro.Combustible,
		"consumo":        NuevoCarro.Consumo,
		"almacenamiento": NuevoCarro.Almacenamiento,
		"descripcion":    NuevoCarro.Descripcion,
		"lujo":           NuevoCarro.Lujo,
		"deportivo":      NuevoCarro.Deportivo,
		"imagen":         NuevoCarro.Imagen,
	}
	NuevoId, err := c.repo.Create(context.Background(), CreateQueryC, valores_columnas)
	if err != nil {
		log.Printf("fallo a crear un carro, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo a crear un carro, con error: %s", err.Error())
	}
	return NuevoId, nil
}

func (c *ControllerU) CrearUsuario(body []byte) (int64, error) {
	NuevoUsuario := &models.Usuario{}
	err := json.Unmarshal(body, NuevoUsuario)
	if err != nil {
		log.Printf("fallo a crear un usuario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo a crear un usuario, con error: %s", err.Error())
	}
	valores_columnas := map[string]any{
		"id":         NuevoUsuario.Id,
		"nombre":     NuevoUsuario.Nombre,
		"contrasena": NuevoUsuario.Contrasena,
		"reservas":   NuevoUsuario.Reservas,
	}
	NuevoId, err := c.repo.Create(context.Background(), CreateQueryU, valores_columnas)
	if err != nil {
		log.Printf("fallo a crear un usuario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo a crear un usuario, con error: %s", err.Error())
	}
	return NuevoId, nil
}

func buildUpdateQuery(columnasActualizar map[string]any) string {
	columnas := []string{}
	for key := range columnasActualizar {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	columnasString := strings.Join(columnas, ",")
	return columnasString
}

func (c *Controller) ActualizarCarro(body []byte, id string) error {
	valoresActualizarBody := make(map[string]any)
	err := json.Unmarshal(body, &valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un carro, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un carro, con error: %s", err.Error())
	}
	if len(valoresActualizarBody) == 0 {
		log.Printf("fallo al actualizar un carro, con error: no hay datos en el body")
		return fmt.Errorf("fallo al actualizar un carro, con error:no hay datos en el body")
	}
	UpdtQueryC := fmt.Sprintf(UpdateQueryC, buildUpdateQuery(valoresActualizarBody))
	valoresActualizarBody["id"] = id
	err = c.repo.Update(context.Background(), UpdtQueryC, valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un carro, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un carro, con error: %s", err.Error())
	}
	return nil
}

func (c *ControllerU) ActualizarUsuario(body []byte, id string) error {
	valoresActualizarBody := make(map[string]any)
	err := json.Unmarshal(body, &valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, con error: %s", err.Error())
	}
	if len(valoresActualizarBody) == 0 {
		log.Printf("fallo al actualizar un usuario, con error: no hay datos en el body")
		return fmt.Errorf("fallo al actualizar un usuario, con error:no hay datos en el body")
	}
	UpdtQueryU := fmt.Sprintf(UpdateQueryU, buildUpdateQuery(valoresActualizarBody))
	valoresActualizarBody["id"] = id
	err = c.repo.Update(context.Background(), UpdtQueryU, valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) EliminarCarro(id string) error {
	err := c.repo.Delete(context.Background(), DeleteQueryC, id)
	if err != nil {
		log.Printf("fallo al Eliminar un carro, con error: %s", err.Error())
		return fmt.Errorf("fallo al Eliminar un carro, con error: %s", err.Error())
	}
	return nil
}

func (c *ControllerU) EliminarUsuario(id string) error {
	err := c.repo.Delete(context.Background(), DeleteQueryU, id)
	if err != nil {
		log.Printf("fallo al Eliminar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al Eliminar un usuario, con error: %s", err.Error())
	}
	return nil
}
