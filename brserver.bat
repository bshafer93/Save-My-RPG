@ECHO OFF
docker build -t savemyrpg-production -f Dockerfile.production .
docker run -it --rm --network smrpg --name SaveMyRpgServer -p 8100:8100 savemyrpg-production