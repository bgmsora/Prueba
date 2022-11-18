# Prueba
Prueba técnica: Data pipeline
Brandon Gómez

## Pasos

### 1. Crear base de datos con GraphQL
1. Se ocupará PostgreSQL como base de datos y para GraphQL se usara Hasura.
2. Se cargarán los datos del csv a la base de datos (https://www.convertcsv.com/csv-to-sql.htm)
3. GraphQL url:http://localhost:8080/console
4. Es necesario entrar en el apartado de Data -> Databases -> default -> public -> Untracked tables or views y hacer click en **Track All**, para rastrear los datos que tiene PostgreSQL (Solo la primera vez que se ejecute docker-compose up)
*Como se muestra en la siguiente imagen:*

![Alt text](doc/resource/track.png?raw=true "Title")

### 2. Analisis de los datos
1. Traduce coordenadas a direcciones ocupando una Api externa, ya que al ocupar un archivo de configuración que contenía todas las coordenadas ocupaba mucha memoria para incluirlo en un docker

### 3. Kubernetes (k8s)
Para el despliegue se tienen los archivos kubemanifests.yaml y su archivo de variables de ambiente env.yaml

### Notas
Todos los datos de acceso o configuraciones se encuentran en el archivo **.env** y **env.yaml**, que viene adjunto en el correo.
Se publico una imagen de la Api para poder ser utilizada en kubernetes en el siguiente enlace:
https://hub.docker.com/repository/docker/bgmrand/test

### Postman Workspace
Pruebas integración
ID:  dea07cb7-98de-49fe-bdc7-d628ecce7c23

### Pruebas unitarias
Estas se encuentran en la carpeta services en el archivo api_test.go, se ejecutan con:
```console
foo@bar:~$ go test
```




### Diagrama de secuencia 
![Alt text](doc/resource/SequenceDiagram.png?raw=true "Title")