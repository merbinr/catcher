FROM golang:1.22-alpine AS build_env
WORKDIR /app
COPY . .
RUN go build -o catcher ./cmd/*.go


FROM alpine:3.20
WORKDIR /app
COPY --from=build_env /app /app
EXPOSE 8080
ENTRYPOINT [ "/app/catcher" ]
