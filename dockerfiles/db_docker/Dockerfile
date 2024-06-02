# Use official MySQL Docker image as base image
FROM mysql:latest

# Environment variables to configure MySQL
ENV MYSQL_DATABASE burp_backend
ENV MYSQL_ROOT_PASSWORD password

# Copy SQL initialization scripts to docker-entrypoint-initdb.d directory
COPY ./init.sql /docker-entrypoint-initdb.d/
