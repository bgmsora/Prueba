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
1. Se revisaran las consultas que se piden y los datos que se tienen para dar una solucion optima
2. Traducir coordenadas a direcciones



### Nota
Todas los datos de acceso o configuraciones se encuentran en el archivo .env, que fue enviado con el correo


### Postman Workspace
Pruebas integrales 
ID:  dea07cb7-98de-49fe-bdc7-d628ecce7c23
