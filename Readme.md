drop database paimonbank;
create database paimonbank;
GRANT ALL ON DATABASE paimonbank TO abdillah;
ALTER DATABASE paimonbank OWNER TO abdillah;

mingw32-make migrate_up

testing on k6
$env:BASE_URL = 'http://localhost:8080'
k6 run --vus 1 --iterations 1 script.js

migrate -path db/migrations -database "postgresql://abdillah:pass@localhost:8000/paimonbank?sslmode=disable" up

aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 560918124458.dkr.ecr.ap-southeast-1.amazonaws.com

docker tag budimanbank:1.0.0 560918124458.dkr.ecr.ap-southeast-1.amazonaws.com/budimanbank:1.0.0

docker push 560918124458.dkr.ecr.ap-southeast-1.amazonaws.com/budimanbank:1.0.0
