
**DB diagram**: link: https://dbdiagram.io/d/63ff1db5296d97641d849361

### Run Postgres Docker image
```bash
docker pull postgres:15.2-alpine
```

```bash
docker run --name my-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2-alpine
```

Test the container:
```bash
docker exec -it my-postgres psql -U root
```