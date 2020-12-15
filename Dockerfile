FROM golang:1.15-alpine AS build
WORKDIR /build
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o excelante .

FROM scratch
WORKDIR /bin
COPY --from=build /build/excelante .
EXPOSE 8000
CMD ["./excelante"]
