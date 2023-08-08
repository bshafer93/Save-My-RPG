docker build -t savemyrpg-production -f Dockerfile.production .

docker run -it --rm --network smrpg -p 8100:8100 savemyrpg-production

docker run -it --rm --network smrpg -p 8200:8200 smrpg_db
docker run -it --rm --network smrpg -p 5432:5432 --name smrpg_db -e POSTGRES_PASSWORD=ninjame -d postgres

db.createUser(
  {
    user: "admin",
    pwd:  passwordPrompt(),   // or cleartext password
    roles: [ { role: "readWrite", db: "smrpg_db" } ]
  }
)

mongodb://admin:ninjame@192.168.1.33:8200

docker run -it --rm --network smrpg postgres psql -U smrpg_db

docker compose -f docker-compose.yaml up

psql -U admin -d default -a -f /db_scripts/db_initialize.sql