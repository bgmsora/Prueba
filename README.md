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
Traducir coordenadas a direcciones ocupando una Api externa, ya que al ocupar un archivo de configuración que contenía todas las coordenadas ocupaba mucha memoria para incluirlo en un docker

### 3. Kubernetes (k8s)
Para el despliegue se tienen los archivos kubemanifests.yaml y su archivo de variables de ambiente env.yaml

### Notas
Todos los datos de acceso o configuraciones se encuentran en el archivo **.env** y **env.yaml**, que viene adjunto en el correo.
Se publico una imagen de la Api para poder ser utilizada en kubernetes en el siguiente enlace:
https://hub.docker.com/repository/docker/bgmsora/test

Si no se está corriendo el servicio de K8s en la nube se recomienda usar minikube, porque al querer hacer la menor cantidad de configuraciones no se hizo el uso de ingress porque se requería configurar externamente con Nginx o en su defecto modificar el archivo  */etc/hosts*
Por lo cual se recomienda lo siguiente:
```console
foo@bar:~$ minikube start
foo@bar:~$ kubectl apply -f kubemanifests.yaml
foo@bar:~$ kubectl apply -f env.yaml
foo@bar:~$ minikube tunnel
```
También en Hasura en la pestaña de Data, en el anexo SQL copiar el archivo 
**/db/init.sql**
Para rellenar la base de datos

### Postman Workspace
Pruebas integración
https://www.postman.com/flight-physicist-92712039/workspace/prueba-tcnica/overview

### Pruebas unitarias
Estas se encuentran en la carpeta services en el archivo api_test.go, se ejecutan con:
```console
foo@bar:~$ go test
```

### Diagrama de secuencia 
![Alt text](doc/resource/SequenceDiagram.png?raw=true "Title")