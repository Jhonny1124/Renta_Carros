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
    imagen TEXT
);

INSERT INTO Carro (marca, referencia, modelo, tipo, potencia, torque, transmision, motor, pasajeros, combustible, consumo, almacenamiento, descripcion, lujo, deportivo, imagen) 
VALUES ('Kia', 'Kia Picanto Zenith', '2020', 'Hatchback', '83 HP', '112 Nm','Manual', '1.25L', '5 Pasajeros', 'Corriente', '70 km/galon', '255L', 'Carro pequeño para transportarse en la ciudad', 0, 0, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTRedPGs70M7UOKozuucEnt-ENxYuEGB_x_Fg&s');

INSERT INTO Carro (marca, referencia, modelo, tipo, potencia, torque, transmision, motor, pasajeros, combustible, consumo, almacenamiento, descripcion, lujo, deportivo, imagen) 
VALUES ('Volkswagen', 'Volkswagen Golf R', '2022', 'Hatchback', '315 HP', '400 Nm','Automatica', '2.0L', '5 Pasajeros', 'Extra', '50 km/galon', '341L', 'Carro con potencia increible y un diseño muy deportivo', 1, 1, 'https://www.namesnack.com/images/Namesnack-nombres-para-concesionarios-de-autos-5472x3648-20230523.jpeg?crop=21:16,smart&width=420&dpr=2');

CREATE TABLE IF NOT EXISTS Usuario (
    id SERIAL PRIMARY KEY,
    nombre TEXT,
    contrasena TEXT,
    reservas TEXT
);

INSERT INTO Usuario (nombre, contrasena, reservas) VALUES ('alejoadmin', 'alejo1234', ' ');