to run a mariadb instance:

```shell
docker run -d --rm --name mariadb -e MARIADB_ROOT_PASSWORD=11111 -e MARIADB_DATABASE=hello -p 3306:3306  mariadb:latest
```
