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

 
async function buscar12() {
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

function MostrarCategorias(){
  document.getElementById('categoria').style.opacity = '1';
  document.getElementById('categoria').style.pointerEvents = 'all';
  document.getElementById('marcas').style.opacity = '0';
  document.getElementById('marcas').style.pointerEvents = 'none';
  document.getElementById('tipos').style.opacity = '0';
  document.getElementById('tipos').style.pointerEvents = 'none';
  document.getElementById('transmision').style.opacity = '0';
  document.getElementById('transmision').style.pointerEvents = 'none';
  document.getElementById('combustible').style.opacity = '0';
  document.getElementById('combustible').style.pointerEvents = 'none';
}

function OcultarCategorias(){
  document.getElementById('categoria').style.opacity = '0';
  document.getElementById('categoria').style.pointerEvents = 'none';
  document.getElementById('marcas').style.opacity = '0';
  document.getElementById('marcas').style.pointerEvents = 'none';
  document.getElementById('tipos').style.opacity = '0';
  document.getElementById('tipos').style.pointerEvents = 'none';
  document.getElementById('transmision').style.opacity = '0';
  document.getElementById('transmision').style.pointerEvents = 'none';
  document.getElementById('combustible').style.opacity = '0';
  document.getElementById('combustible').style.pointerEvents = 'none';
}

function MostrarMarcas(){
  document.getElementById('categoria').style.opacity = '0';
  document.getElementById('categoria').style.pointerEvents = 'none';
  document.getElementById('marcas').style.opacity = '1';
  document.getElementById('marcas').style.pointerEvents = 'all';
  document.getElementById('tipos').style.opacity = '0';
  document.getElementById('tipos').style.pointerEvents = 'none';
  document.getElementById('transmision').style.opacity = '0';
  document.getElementById('transmision').style.pointerEvents = 'none';
  document.getElementById('combustible').style.opacity = '0';
  document.getElementById('combustible').style.pointerEvents = 'none';
}

function MostrarTipos(){
  document.getElementById('categoria').style.opacity = '0';
  document.getElementById('categoria').style.pointerEvents = 'none';
  document.getElementById('marcas').style.opacity = '0';
  document.getElementById('marcas').style.pointerEvents = 'none';
  document.getElementById('tipos').style.opacity = '1';
  document.getElementById('tipos').style.pointerEvents = 'all';
  document.getElementById('transmision').style.opacity = '0';
  document.getElementById('transmision').style.pointerEvents = 'none';
  document.getElementById('combustible').style.opacity = '0';
  document.getElementById('combustible').style.pointerEvents = 'none';
}

function MostrarTransmision(){
  document.getElementById('categoria').style.opacity = '0';
  document.getElementById('categoria').style.pointerEvents = 'none';
  document.getElementById('marcas').style.opacity = '0';
  document.getElementById('marcas').style.pointerEvents = 'none';
  document.getElementById('tipos').style.opacity = '0';
  document.getElementById('tipos').style.pointerEvents = 'none';
  document.getElementById('transmision').style.opacity = '1';
  document.getElementById('transmision').style.pointerEvents = 'all';
  document.getElementById('combustible').style.opacity = '0';
  document.getElementById('combustible').style.pointerEvents = 'none';
}

function MostrarCombustible(){
  document.getElementById('categoria').style.opacity = '0';
  document.getElementById('categoria').style.pointerEvents = 'none';
  document.getElementById('marcas').style.opacity = '0';
  document.getElementById('marcas').style.pointerEvents = 'none';
  document.getElementById('tipos').style.opacity = '0';
  document.getElementById('tipos').style.pointerEvents = 'none';
  document.getElementById('transmision').style.opacity = '0';
  document.getElementById('transmision').style.pointerEvents = 'none';
  document.getElementById('combustible').style.opacity = '1';
  document.getElementById('combustible').style.pointerEvents = 'all';
}

const contenido = document.getElementById("contenido");

let matriz = document.createElement("div");
matriz.classList.add("productos");

let iniciopagina = document.getElementById("contenidoinicial");

let PaginasMostradas = document.createElement("div");
PaginasMostradas.style.justifyContent = "center";
PaginasMostradas.style.textAlign = "center";



let carro = document.createElement("div");

let apiurl;
let datos;
let nombre_carro;



function tipoFiltro(valor){//Funcion que determina los links dpendiendo del tipo de busquedad
  switch(valor){
      case 1:
          apiurl = "http://localhost:8080/cars";
          break;
      case 2:
          apiurl = "http://localhost:8080/cars/marca/";
          break;
      case 3:
          apiurl = "http://localhost:8080/cars/tipo/";
          break;
      case 4:
          apiurl = "http://localhost:8080/cars/transmision/";
          break;
      case 5:
          apiurl = "http://localhost:8080/cars/combustible/";
          break;
  }
};

async function buscar(entrada){//Funcion que evalua si la busquedad esta correcta y la ejecuta
  const url = apiurl+entrada;
  console.log(url)
  try{
    const respuesta = await getdata(url);
    if(respuesta == null){
      Swal.fire({
        title: 'Sin resultados',
        text: 'Lo sentimos, no se encontraron resultados para la búsqueda',
        icon: 'info',
        confirmButtonText: 'OK'
      });
    }
    else{
      productos(respuesta);
    }
  }
  catch(error){
    console.error(`fallo la consulta a la api: ${error}`);
    return
  }
}

function productos(datos) {//Funcion que crea tarjetas para los carros
  carro.innerHTML = "";
  matriz.innerHTML = "";
  iniciopagina.style.display = "none";
  PaginasMostradas.innerHTML = "";
  let ResultadosEncontrados = document.createElement("h2");
  ResultadosEncontrados.textContent = "Results found ("+datos.length+")";
  ResultadosEncontrados.classList.add("encabezado2");

  PaginasMostradas.appendChild(ResultadosEncontrados);

  // Determinar el número total de páginas
  const totalPaginas = Math.ceil(datos.length / 9);

  // Función para mostrar una página específica
  function mostrarPagina(pagina) {
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    });
    matriz.innerHTML = ""; // Limpiar el contenido actual de la matriz
    const inicio = pagina * 9;
    const fin = Math.min(inicio + 9, datos.length);

    for (let i = inicio; i < fin; i++) {
      const dato = datos[i];

      let tarjeta = document.createElement("a");
      tarjeta.classList.add("tarjeta");
      tarjeta.addEventListener("click", function(){
        /*info_producto(tarjeta);*/
      });

      let nombre= document.createElement("p");
      nombre.textContent = dato.referencia;
      nombre.classList.add("prueba");

      let imagen = document.createElement("img");
      imagen.classList.add("imagen-tarjeta");
      imagen.src = dato.imagen;
      imagen.alt = dato.referencia;

      tarjeta.addEventListener("mouseenter", function(){
        nombre.style.fontSize = "min(3.5vw, 37px)";
        nombre.style.transition = "font-size 0.4s ease"; 
        imagen.style.opacity = "0.65";
        imagen.style.transition = "opacity 0.4s ease"; 
      });
      tarjeta.addEventListener("mouseleave", function(){
        nombre.style.fontSize = "min(3vw, 33px)";
        nombre.style.transition = "font-size 0.4s ease"; 
        imagen.style.opacity = "0.75";
        imagen.style.transition = "opacity 0.4s ease"; 
      });

      let boton = document.createElement("button");
      boton.textContent = "Reservar";

      let precio = document.createElement("p");
      precio.textContent = dato.precio;

      tarjeta.appendChild(imagen);
      tarjeta.appendChild(nombre);
      tarjeta.appendChild(boton);
      tarjeta.appendChild(precio);

      matriz.appendChild(tarjeta);
    }
  }

  // Mostrar la primera página al cargar
  mostrarPagina(0);

  // Agregar botones para navegar entre las páginas
  PaginasMostradas.appendChild(matriz);
  for (let i = 0; i < totalPaginas; i++) {
    const boton = document.createElement("button");
    boton.classList.add("BotonPagina");
    boton.textContent = i + 1; // El índice de la página comienza desde 0
    boton.addEventListener("click", function() {
      mostrarPagina(i);
    });
    PaginasMostradas.appendChild(boton);
  }
  contenido.appendChild(PaginasMostradas);
}

function Inicio(){
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  });
  carro.innerHTML = "";
  matriz.innerHTML = "";
  iniciopagina.style.display = "block";
  PaginasMostradas.innerHTML = "";
}