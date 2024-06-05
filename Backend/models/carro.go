package models

type Carro struct {
	Id             int    `json:"id" db:"id"`
	Marca          string `json:"marca" db:"marca"`
	Referencia     string `json:"referencia" db:"referencia"`
	Modelo         string `json:"modelo" db:"modelo"`
	Tipo           string `json:"tipo" db:"tipo"`
	Potencia       string `json:"potencia" db:"potencia"`
	Torque         string `json:"torque" db:"torque"`
	Transmision    string `json:"transmision" db:"transmision"`
	Motor          string `json:"motor" db:"motor"`
	Pasajeros      string `json:"pasajeros" db:"pasajeros"`
	Combustible    string `json:"combustible" db:"combustible"`
	Consumo        string `json:"consumo" db:"consumo"`
	Almacenamiento string `json:"almacenamiento" db:"almacenamiento"`
	Descripcion    string `json:"descripcion" db:"descripcion"`
	Lujo           int    `json:"lujo" db:"lujo"`
	Deportivo      int    `json:"deportivo" db:"deportivo"`
	Imagen         string `json:"imagen" db:"imagen"` /*URL de la imagen*/
	Reservado      int    `json:"reservado" db:"reservado"`
}
