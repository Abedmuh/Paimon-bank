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

migrate -path ./db/migrations -database "postgres://p_tangoredox:iFaph6iarahBahcuethee0Ue7nee9ejie@projectsprint-db.cavsdeuj9ixh.ap-southeast-1.rds.amazonaws.com:5432/tangoredox?sslmode=verify-full&sslrootcert=ap-southeast-1-bundle.pem" up
