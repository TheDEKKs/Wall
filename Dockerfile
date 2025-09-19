FROM golang:1.24.4
WORKDIR /app
COPY .. /app

ENTRYPOINT ["go", "mod", "tidy"]
ENTRYPOINT ["go", "run", "./cmd/partywall"]
