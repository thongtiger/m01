# builder
# FROM golang:alpine as builder
FROM golang:1.12.5-stretch as builder
ENV GO111MODULE=on

WORKDIR /go-modules/m01

# Get dependancies - will also be cached if we won't change mod/sum
# COPY the source code as the last step
COPY  . .

# Build the binary
# RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod=vendor -o m01 .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o m01 .
# RUN CGO_ENABLED=0 GOOS=linux go build -o m01 .


# # Runner 1
# FROM alpine:3.8
# WORKDIR /root/
# COPY --from=builder /go-modules/m01 .

# # env  
# ENV PORT $PORT
# ENV JWT_PUBLIC_KEY $JWT_PUBLIC_KEY
# # mssql env
# ENV MSSQL_HOST $MSSQL_HOST
# ENV MSSQL_DB $MSSQL_DB
# ENV MSSQL_USERNAME $MSSQL_USERNAME
# ENV MSSQL_PASSWORD $MSSQL_PASSWORD
# # mongo env
# ENV MONGO_HOST $MONGO_HOST
# ENV MONGO_DB $MONGO_DB
# ENV MONGO_USER $MONGO_USER
# ENV MONGO_PASSWORD $MONGO_PASSWORD

# EXPOSE $PORT
# CMD ["./m01"]



# Runner 2
FROM busybox
ARG PORT
COPY --from=builder /go-modules/m01 .

# env from kubernetes
ENV PORT $PORT
ENV JWT_PUBLIC_KEY $JWT_PUBLIC_KEY
# mssql env
ENV MSSQL_HOST $MSSQL_HOST
ENV MSSQL_DB $MSSQL_DB
ENV MSSQL_USER $MSSQL_USER
ENV MSSQL_PW $MSSQL_PW
# mongo env
ENV MONGO_HOST $MONGO_HOST
ENV MONGO_DB $MONGO_DB
ENV MONGO_USER $MONGO_USER
ENV MONGO_PW $MONGO_PW

USER root
EXPOSE $PORT
CMD ["./m01"]
