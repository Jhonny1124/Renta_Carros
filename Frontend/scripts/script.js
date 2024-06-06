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

async function IniciarSesion(){
  const datos = await getdata('http://localhost:8080/users');
  const nombreBuscado = document.getElementById('nombreusuario').value;
  sessionStorage.setItem('nombreUsuario', nombreBuscado)
  const usuarioencontrado = datos.find(usuario => usuario.nombre === nombreBuscado)
  if (usuarioencontrado === undefined){
    Swal.fire({
      title: 'Error',
      text: 'No se pudo iniciar sesión. Verifica tu nombre de usuario y contraseña.',
      icon: 'error',
      confirmButtonText: 'OK'
    });
    return
  }
  else{
    if(usuarioencontrado.contrasena != document.getElementById('contrasena').value){
      Swal.fire({
        title: 'Error',
        text: 'No se pudo iniciar sesión. Verifica tu nombre de usuario y contraseña.',
        icon: 'error',
        confirmButtonText: 'OK'
      });
      return
    }
    else{
      sessionStorage.setItem('contrasenaUsuario', document.getElementById('contrasena').value)
      document.getElementById('nombreusuario').value='';
      window.location.href ='../paginas/paginainicio.html'
    }
  }
}

async function CrearUsuario(){

  const datos = await getdata('http://localhost:8080/users');

  const nombreBuscado = document.getElementById('nombreusuarionuevo').value;
  const contrasena = document.getElementById('contrasenanueva').value;
  const confirmacion = document.getElementById('confirmarcontrasena').value;

  const usuarioencontrado = datos.find(usuario => usuario.nombre === nombreBuscado)

  if(nombreBuscado.replaceAll(' ','') === ''){
    Swal.fire({
      title: 'Error',
      text: 'Ingresa un nombre de usuario valido',
      icon: 'error',
      confirmButtonText: 'OK'
    });
    return
  }
  else if(contrasena.replaceAll(' ','') === '' || confirmacion.replaceAll(' ','') === ''){
    Swal.fire({
      title: 'Error',
      text: 'Ingresa una contraseña de usuario valida',
      icon: 'error',
      confirmButtonText: 'OK'
    });
    return
  }

  if (usuarioencontrado === undefined){
    if(contrasena.length < 6){
      Swal.fire({
        title: 'Error',
        text: 'La contraseña es demasiado corta. Minimo 6 caracteres',
        icon: 'error',
        confirmButtonText: 'OK'
      });
      return
    }
    if(contrasena === confirmacion){
      const playload = {
        nombre: nombreBuscado,
        contrasena: contrasena,
        reservas: ''
      };
      postdata('http://localhost:8080/users',playload)
      Swal.fire({
        title: 'Exito',
        text: 'Usuario Creado Correctamente',
        icon: 'success',
        confirmButtonText: 'OK'
      });
      PasarInicioSesion()
      return
    }
    else{
      Swal.fire({
        title: 'Error',
        text: 'Las contraseñas no coinciden. Por favor, verifica y vuelve a intentarlo.',
        icon: 'error',
        confirmButtonText: 'OK'
      });
      return
    }
  }
  else{
    Swal.fire({
      title: 'Error',
      text: 'El nombre de usuario ya está en uso. Por favor, elige otro nombre de usuario.',
      icon: 'error',
      confirmButtonText: 'OK'
    });
    return
  }
}

function PasarCrearUsuario(){
  document.getElementById('iniciosesion').style.opacity = '0';
  document.getElementById('iniciosesion').style.pointerEvents = 'none';
  document.getElementById('CrearUsuario').style.opacity = '1';
  document.getElementById('CrearUsuario').style.pointerEvents = 'all';

  document.getElementById('nombreusuario').value='';
  document.getElementById('nombreusuarionuevo').value='';
  document.getElementById('contrasena').value='';
  document.getElementById('contrasenanueva').value='';
  document.getElementById('confirmarcontrasena').value='';
}

function PasarInicioSesion(){
  document.getElementById('iniciosesion').style.opacity = '1';
  document.getElementById('iniciosesion').style.pointerEvents = 'all';
  document.getElementById('CrearUsuario').style.opacity = '0';
  document.getElementById('CrearUsuario').style.pointerEvents = 'none';
}

function VerificacionEnter(evento,eleccion){//Funcion que permite hacer busquedad con enter
  if(evento.keyCode === 13){
    if(eleccion == 0){
      IniciarSesion();
    }
    else{
      CrearUsuario();
    }
  }
}