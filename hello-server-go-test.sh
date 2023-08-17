chmod +x ./linux/hello-server

./linux/hello-server &

sleep 3

for i in Person1 SecondPerson Me Ender;
do
  echo "$(date): $(curl -s http://localhost:8080/${i})"
done
