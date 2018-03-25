# gin-web-example

Gin Web Example

To execute the application from the command line, go to your application directory and execute the following command:

```
go build -o app
```

This will build your application and create an executable named app which you can run as follows:

```
./app
```

# Docker

This will build docker image

```
docker build -t gin-web-example .
```

This will run docker image on port 80

```
docker run -d -p 80:8080 gin-web-example
```
