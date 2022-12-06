# sooa-user-token-ms-dos

Registro y autenticación basados en tokens para usuarios de SOOA

docker build --tag degarzonm/sooa-user-token-ms-dos .
docker tag degarzonm/sooa-user-token-ms-dos degarzonm/sooa-user-token-ms-dos:v1.4
docker push degarzonm/sooa-user-token-ms-dos:v1.4
docker run -p 6663:6663 degarzonm/sooa-user-token-ms-dos -e PORT=6664 MONGO_URI=mongodb+srv://sooa_mongo_admin:4MbzZzD9ykOMrvLo@sooa-mongo-cluster.lrlq0px.mongodb.net/?retryWrites=true&w=majority

cloud

docker pull degarzonm/sooa-user-token-ms:v1.4