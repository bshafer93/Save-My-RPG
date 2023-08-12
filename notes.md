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

docker run -d --network smrpg cloudflare/cloudflared:latest tunnel --no-autoupdate run --token eyJhIjoiZGQyNDcxNTAwODU2NTE5NWRhZmFkN2ZjMzlkNTAwYjciLCJ0IjoiNTRiMTM5NWQtZTE2NC00Njc1LTg2NTUtNzJhYjhmODU4ODZhIiwicyI6Ik5HTmlOVGM0T0RRdFlqTm1aUzAwTUdGbUxXSTNOR1l0WkdJM00yWTNNbVpsTnpReSJ9

https://www.youtube.com/watch?v=QXooywQSfJY - set up docker in truecharts 


docker run -d --name=tailscaled --network smrpg --cap-add=NET_ADMIN --socks5-server=localhost:8100 --cap-add=NET_RAW TS_AUTHKEY=tskey-auth-ab1CDE2CNTRL-0123456789abcdef tailscale/tailscale