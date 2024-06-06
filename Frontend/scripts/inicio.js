async function getdata(url) {
    //Funcion para consultar la API
      try {
        const response = await axios.get(url);
        const data = response.data;
        return data;
      } catch (error) {
        console.error(`fallo la consulta a la api: ${error}`);
      }
}

async function postdata(url, playload) {
    //Funcion para consultar la API
      try {
        const response = await axios.post(url, playload);
        const data = response.data;
        return data;
      } catch (error) {
        console.error(`fallo la consulta a la api: ${error}`);
      }
}

async function updatedata(url, playload) {
    //Funcion para consultar la API
      try {
        const response = await axios.patch(url, playload);
        const data = response.data;
        return data;
      } catch (error) {
        console.error(`fallo la consulta a la api: ${error}`);
      }
}

async function deletedata(url) {
    //Funcion para consultar la API
      try {
        const response = await axios.delete(url);
        const data = response.data;
        return data;
      } catch (error) {
        console.error(`fallo la consulta a la api: ${error}`);
      }
}

 
async function buscar() {
    const playload ={
		marca:          'Toyota',
		referencia:     'Toyota TXL',
		modelo:         '2022',
		tipo:           'SUV',
		potencia:       '271 HP a 5600 rpm',
		torque:         '410 Nm a 2800 rpm',
		transmision:    'Automatica',
		motor:          '3.9L',
		pasajeros:      '8 pasajeros',
		combustible:    'Diesel',
		consumo:        '31 km/galon',
		almacenamiento: '840 Kg',
		descripcion:    'Camioneta de lujo que te permite llevar a toda tu familia a una gran avanetura',
		lujo:           1,
		deportivo:      0,
		imagen:         'https://www.elcarrocolombiano.com/wp-content/uploads/2022/06/20220623-COMPRAR-CARRO-NUEVO-POR-IMPORTACION-DIRECTA-VENTAJAS-Y-DESVENTAJAS-01.jpg',
		reservado:     0
    };

    const playload2 ={
		referencia:     'Toyota Prado TXL',
		modelo:         '2023',
    };
    const datos = await getdata('http://localhost:8080/cars');
    console.log(datos);
}

const nombreUsuario = sessionStorage.getItem('nombreUsuario');
const contrasena = sessionStorage.getItem('contrasenaUsuario');

function prueba(){
    console.log(nombreUsuario)
    console.log(contrasena)
}