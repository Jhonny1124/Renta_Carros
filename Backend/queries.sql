CREATE TABLE IF NOT EXISTS Carro (
    id SERIAL PRIMARY KEY,
    marca TEXT,
    referencia TEXT,
    modelo TEXT,
    tipo TEXT,
    potencia TEXT,
    torque TEXT,
    transmision TEXT,
    motor TEXT,
    pasajeros TEXT,
    combustible TEXT,
    consumo TEXT,
    almacenamiento TEXT,
    descripcion TEXT,
    lujo INT,
    deportivo INT,
    imagen TEXT,
    reservado INT,
    precio TEXT
);

INSERT INTO Carro (marca, referencia, modelo, tipo, potencia, torque, transmision, motor, pasajeros, combustible, consumo, almacenamiento, descripcion, lujo, deportivo, imagen, reservado, precio) 
VALUES ('Kia', 'Kia Picanto Zenith', '2020', 'Hatchback', '83 HP', '112 Nm','Manual', '1.25L', '5 Pasajeros', 'Corriente', '70 km/galon', '255L', 'Carro pequeño para transportarse en la ciudad', 0, 0, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTRedPGs70M7UOKozuucEnt-ENxYuEGB_x_Fg&s',0,'$90.000/Dia');

INSERT INTO Carro (marca, referencia, modelo, tipo, potencia, torque, transmision, motor, pasajeros, combustible, consumo, almacenamiento, descripcion, lujo, deportivo, imagen, reservado, precio) 
VALUES ('Volkswagen', 'Volkswagen Golf R', '2022', 'Hatchback', '315 HP', '400 Nm','Automatica', '2.0L', '5 Pasajeros', 'Extra', '50 km/galon', '341L', 'Carro con potencia increible y un diseño muy deportivo', 1, 1, 'https://www.elcarrocolombiano.com/wp-content/uploads/2020/11/2022_Golf_R_European_model_shown-Large-12436.jpg',0,'$300.000/Dia');

CREATE TABLE IF NOT EXISTS Usuario (
    id SERIAL PRIMARY KEY,
    nombre TEXT,
    contrasena TEXT,
    reservas TEXT
);

INSERT INTO Usuario (nombre, contrasena, reservas) VALUES ('alejoadmin', 'alejo1234', ' ');
