Se construye los containers con ```docker compose up --build```

El back se encuentra en el puerto ```8000```, el db en ```3306``` y el front se accede desde ```3000``` (aunque en el container se encuentra en el puerto 80 por ngix).

La vista de react se encuentra en ```http://localhost:3000/```