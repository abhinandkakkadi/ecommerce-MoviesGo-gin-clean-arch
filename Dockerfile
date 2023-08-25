
# stage 1
FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api

EXPOSE 3000

CMD ["/api"]

FROM alpine:3.14 AS prod

# # Set the working directory
WORKDIR /app

# Copy only the necessary files from the build image
COPY --from=build /api /api
COPY --from=build . ./

# expose container port
EXPOSE 3000

# Set the entry point
CMD ["/api"]
