```bash
url -X POST http://localhost:8080/albums \
-H "Content-Type: application/json" \
-d '{
  "title": "Kind of Blue",
  "artist": "Miles Davis",
  "price": 29.99
}'
```

```bash
curl http://localhost:8080/albums/5
```

- `\d` describe table;
- `\?` list all the commands
- `\l` list databases
- `\conninfo` display information about current connection
- `\c [DBNAME]` connect to new database, e.g., \c template1
- `\dt` list tables of the public schema
- `\dt <schema-name>.*` list tables of certain schema, e.g., \dt public.*
- `\dt *.*` list tables of all schemas
- `\q` quit psql