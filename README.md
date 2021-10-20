# Api Products


### Descripcion

 Api-Products expone el tipico crud sobre el dominio de products.

 - Crear 
 - Editar
 - Eliminar
 - Actualizar
 - Listar

 Se responden con los clasicos http codes, consumiendo y respondiendo JSON

Se separa las responsabilidades de la api por paquetes e interfaces: 
- Internal (logica de negocio) 
- Pkg interfaces externas

Esto para facilitar la extension y/o modificacion de las implementaciones, ademas de los testing unitarios.

### Conceptos

- Inyeccion de dependencias
- Unica instancia de bd
- Cerrar conexion a la bd
- Uso de decoradores
- Composicion
- Test unitarios
- Mocks 
- Separacion de responsabilidades

### TODO: Implement
- Implementar api-key y endpoint de autenticacion por token
- Agregar test a la validacion del modelo Product
- Mejorar manejo de errores en la capa de servicio
- Agregar script que espere a iniciar el contenedor de la api cuando el servidor mongo ya este disponible
- Cubrir mas casos de pruebas
- Agregar ci/cd
### Instrucciones

- Instalar go 1.17
- Instalar mongo o configurar archivo .env hacia un servidor mongo
- Crear base de datos bd y coleccion products
- Configurar demas variables de entorno
- Importar archivo de coleccion con Postman 
- Desde el terminal ejecutar go run cmd/api/main.go
- Ejecutar requests

### Test & Coverage

- go test ./... -cover  
- go test tool cover -func profile.cov
- go tool cover -html profile.cov -o coverage.html