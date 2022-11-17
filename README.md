# Prueba
Prueba técnica: Data pipeline

## Pasos

### 1. Se creara la base de datos con GraphQL
1. Se ocupara PostgreSQL como base de datos y para GraphQL se usara Hasura.
2. Se cargaran los datos del csv a la base de datos (https://www.convertcsv.com/csv-to-sql.htm)
3. GraphQL http://localhost:8080/console
4. Es necesario entrar en el apartado de Data, Databases, default, public, Untracked tables or views y presionar Track All, para rastrear todos los datos que tiene la base de datos de PostgreSQL (Solo la primer vez que se ejecute docker-compose up)

![Alt text](doc/resource/track.png?raw=true "Title")

### 2. Analisis de los datos
1. Para traducir coordenadas a direcciones, se ocupo una Api externa, porque se intento ocupar un json interno, pero pesaba demasiado para incluirlo en un docker

### 3. K8s
Para este caso se tiene kubemanifests.yaml y su archivo de enviroment env.yaml

### Nota
Todas los datos de acceso o configuraciones se encuentran en el archivo .env y env.yaml, que fue enviado con el correo

### Postman Workspace
Pruebas integrales 
ID:  dea07cb7-98de-49fe-bdc7-d628ecce7c23

### Pruebas unitarias
Estas se encuentran en la carpetas services en el archivo api_test.go, se puede ejecutar con go test