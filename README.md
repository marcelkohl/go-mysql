# MySQL with Go
sample on using MySQL with Go Language

## The database on a docker
based on https://javabydeveloper.com/create-mysql-container-using-docker-and-docker-compose/

### bringing mysql
```
docker-compose up -d
```
### accessing mysql on the docker
```
docker exec -it mysql-on-docker mysql -uadmin -p
```

### Removing mysql container
```
docker stop mysql-on-docker
docker container rm mysql-on-docker
docker image rm mysql/mysql-server:latest
```
