# Api Golang con Gin-Gonic
#### Seminario Golang, Tudai, Exactas Unicen. Noviembre 2020.
#### Refactorización sobre modelo de Juan Pablo Pizarro.
## Correr api: 
go run cmd/vinoteca/vinotecasrv.go -config ./config/config.yaml
Servidor: localhost:8080
## Rutas:
Get: vinos (trae colección de productos)
Get por ID: vinos/:id (trae producto por identificador único)
Post: postvino (crea un producto, id automático)
Put: putvino/:id (actualiza producto por identificador)
Delete: deletevino/:id (elimina producto por identificador)
