drop database paimonbank;
create database paimonbank;
GRANT ALL ON DATABASE paimonbank TO abdillah;
ALTER DATABASE paimonbank OWNER TO abdillah;

mingw32-make migrate_up

testing on k6
$env:BASE_URL = 'http://localhost:8080'
k6 run --vus 1 --iterations 1 script.js
x
migrate -path db/migrations -database "postgresql://abdillah:pass@localhost:5432/paimonbank?sslmode=disable" up
migrate -path db/migrations -database "postgresql://abdillah:pass@localhost:5432/paimonbank?sslmode=disable" up
