cd frontend

npm run build

cd ..

docker-compose -f docker-compose.yaml up -d --build
