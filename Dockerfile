#go build
FROM golang:1.14-stretch as gobuild
WORKDIR /go/src/app
ADD go.mod .
ADD go.sum .
ADD main.go .
ADD cmd ./cmd
ADD backend ./backend
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cerberus-api .

#angular build
FROM node as ngbuild
WORKDIR /app
ADD . /app
ARG configuration=production
RUN npm install
RUN npm run build -- --outputPath=./dist/out --configuration $configuration

#run container
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=gobuild /go/src/app/cerberus-api /root/cerberus-api
COPY --from=ngbuild /app/dist/out /root/dist/cerberus
COPY ./backend/templates /root/backend/templates
EXPOSE 9090
CMD ["./cerberus-api", "serve"]
