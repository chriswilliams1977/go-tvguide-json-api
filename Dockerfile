#FROM alpine
#EXPOSE 5000
#ADD tvguide /bin/sh
#ENTRYPOINT ["./bin/sh"]

#base image
FROM golang:1.12 as build
#container working dir
WORKDIR /go/src/app
#copy local root to working dir
COPY . .
#enable GO modules
ENV GO111MODULE on
#build go binary
#-o write the resulting executable or object to the named output file
#-v print the names of packages as they are compiled.
RUN go build -v -o /app .

# Now copy it into our base image.
#"Distroless" images contain only your application and its runtime dependencies. 
#They do not contain package managers, shells or any other programs you would expect to find in a standard Linux distribution.
#Restricting what's in your runtime container to precisely what's necessary for your app is a best practice
FROM gcr.io/distroless/base
#copy build output to app
COPY --from=build /app /app
CMD ["/app"]
