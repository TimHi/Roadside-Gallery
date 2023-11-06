docker build -t gallery-frontend .
docker run -it -p 8081:80 --rm --name gallery-frontend gallery-frontend