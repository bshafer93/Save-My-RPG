docker build -t norbertle/savemyrpg -f Dockerfile.production .

docker run -it --rm --network smrpg -p 8100:8100 savemyrpg-production

docker run -it --rm --network smrpg -p 8200:8200 smrpg_db
docker run -it --rm --network smrpg -p 5432:5432 --name smrpg_db -e 

POSTGRES_PASSWORD=ninjame -d postgres


docker run -it --rm --network smrpg postgres psql -U smrpg_db

docker compose -f docker-compose.yaml up

psql -U admin -d default -a -f /db_scripts/db_initialize.sql


ny.storage.bunnycdn.com

https://fonts.google.com/icons?selected=Material+Symbols+Rounded:person:FILL@0;wght@400;GRAD@0;opsz@20&icon.query=user&icon.style=Rounded&icon.platform=web



To detach the tty without exiting the shell, use the escape sequence CTRL+P followed by CTRL+Q. More details here.