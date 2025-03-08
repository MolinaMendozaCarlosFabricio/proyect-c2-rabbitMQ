Proyecti MQInventor-i
Sistema de gestión de pedidos e inventario
Un sistema que gestiona los pedidos, y verifica en un inventario si se
cuenca con existencias de lo pedido
Se requieren de estas entidades:
Producto:
    -ID
    -Nombre
    -Categoría
    -Precio
    -Existencias

Usuario:
    -ID
    -Nombre
    -Apellido
    -Correo
    -Teléfono
    -Contraseña

Pedido:
    -ID
    -Fecha de pedido
    -ID de usuarios
    -ID del estatus

Estatus de Pedido
    -ID
    -Nombre

Adquiere (Relación):
    -ID del pedido
    -ID del producto

1ra API: Se encarga del registro de los pedidos y poner qué productos pide, así como
    usuarios.
2da API: Se encarga de los productos, y decide si cancela el pedido o procede.


Proyecto Lib-ry
Sistema de gestión de biblioteca

Un sistema que gestiona los libros y prestamos de una biblioteca
Para este proceso de negocio, se requiere de estas entidades:
libro:
    -ID
    -Nombre
    -Autor
    -Categoria
    -Editorial
    -Edición
    -Estatus

Categorías:
    -ID
    -Nombre

Usuario:
    -ID
    -Nombre
    -Apellidos
    -INE
    -Deuda

Y siendo esta las entidades a partir de las relaciones:
Prestamo:
    -ID
    -Libro prestado
    -Usuario al que se prestó
    -Fecha de prestamo
    -Fecha de devolución

El proyecto contará con 2 API 's
La primera API, manejará el control de libros y usuarios, 
es decir, agregar, editar y eliminar, y a su vez, al realizar un
prestamo, marcará la información necesaria.

La segunda API, maneja la exposición de registros# proyect-c2-rabbitMQ
