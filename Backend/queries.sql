CREATE TABLE Tipo (
    id INT PRIMARY KEY AUTO_INCREMENT,
    marca VARCHAR(255),
    referencia VARCHAR(255),
    modelo VARCHAR(100),
    carroceria VARCHAR(255)
);
CREATE TABLE Especificaciones(
    id INT PRIMARY KEY AUTO_INCREMENT,
    potencia VARCHAR(100),
    torque VARCHAR(100),
    transmision VARCHAR(100),
    motor VARCHAR(100),
    pasajeros VARCHAR(100),
    combustible VARCHAR(100),
    consumo VARCHAR(100),
    almacenamiento VARCHAR(100)
);
CREATE TABLE InfoAdicional (
    id INT PRIMARY KEY AUTO_INCREMENT,
    descripcion TEXT,
    lujo INT,
    deportivo INT,
    imagen TEXT
);
CREATE TABLE Carro (
    id INT PRIMARY KEY AUTO_INCREMENT,
    tipo_id INT,
    especificaciones_id INT,
    infoadicional_id INT,
    FOREIGN KEY (tipo_id) REFERENCES Tipo(id) ON DELETE CASCADE,
    FOREIGN KEY (especificaciones_id) REFERENCES Especificaciones(id) ON DELETE CASCADE,
    FOREIGN KEY (infoadicional_id) REFERENCES InfoAdicional(id) ON DELETE CASCADE
);