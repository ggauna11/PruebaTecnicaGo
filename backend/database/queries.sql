

CREATE TABLE `Tareas2`
(
    id INT(11) PRIMARY KEY auto_increment,
    tarea varchar(50) NOT NULL,
    descripcion varchar(50) NOT NULL,
    estado varchar(50) NOT NULL,
);

INSERT INTO 
    `clientes` 
    (
       `tarea`,
       `descripcion`,
       `estado`
    )
    VALUES (
      'Limpiar Aulas',
      'Limpiar todas las aulas',
      'Por hacer'
    );
