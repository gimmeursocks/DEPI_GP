FROM alpine:latest

# Install lighttpd web server
RUN apk add --no-cache lighttpd

# Copy files to the default lighttpd document root
COPY . /var/www/localhost/htdocs

# Start lighttpd in foreground mode (-D) with the default configuration
CMD ["lighttpd", "-D", "-f", "/etc/lighttpd/lighttpd.conf"]