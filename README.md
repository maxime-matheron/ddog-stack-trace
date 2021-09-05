# ddog-stack-trace

- Export your DD api key as en env variable
```
export DD_API_KEY="<YOUR_API_KEY>"
```

- Start the DD agent as well as the different servers
```
docker-compose up --build -d
```

- Each server exposes the `/generate-stack` endpoint, target the language you want to targeting `localhost` on the right port:
    - `5000` for Python
    - `5001` for Golang
    - `5002` for Java
    - etc