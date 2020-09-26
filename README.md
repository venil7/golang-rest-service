docker run -d \
 --name some-postgres \
 -e POSTGRES_PASSWORD=mysecretpassword \
 -e PGDATA=/var/lib/postgresql/data/pgdata \
 -v `echo $HOME`/Sources/go/web/data:/var/lib/postgresql/data \
 -p 5432:5432 \
 postgres

```SQL
CREATE TABLE users (
  userid SERIAL PRIMARY KEY,
  name TEXT,
  age INT,
  location TEXT
);
```
