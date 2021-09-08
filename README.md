# ddog-stack-trace

- Export your DD api key as en env variable
```
export DD_API_KEY="<YOUR_API_KEY>"
```

- Start the DD agent as well as the different servers
```
docker-compose up --build -d
```

- Each server exposes the `/generate-stack` endpoint, target the language you want by visiting `localhost:<PORT>` where:
    - `5000` is for Python
    - `5001` is for Golang
    - `5002` is for Java

- You should see the `<Language> stack trace generated!` response in your browser

- Visit the [DD trace explorer](https://app.datadoghq.com/apm/traces) to see the generated traces containing stack traces