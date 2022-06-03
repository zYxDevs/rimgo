FROM --platform=$BUILDPLATFORM golang:alpine AS build

ARG TARGETARCH

WORKDIR /src
RUN apk --no-cache add ca-certificates
COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=$TARGETARCH CGO_ENABLED=0 go build

FROM scratch as bin

WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/rimgo .

EXPOSE 3000

CMD ["/app/rimgo"]
