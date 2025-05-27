FROM golang:1.24-bullseye AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/godocker

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/godocker /
COPY ./static ./static
EXPOSE 8080
CMD ["/godocker"]
