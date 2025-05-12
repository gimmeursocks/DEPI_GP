FROM alpine:latest

# Install lighttpd web server and envsubst (from gettext)
RUN apk add --no-cache lighttpd gettext

# Copy files to the default lighttpd document root
COPY . /var/www/localhost/htdocs

# Copy the entrypoint script & make it executable
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Expose the app port
EXPOSE 80

# Start lighttpd in foreground mode (-D) with the default configuration
# CMD ["lighttpd", "-D", "-f", "/etc/lighttpd/lighttpd.conf"]

# Use custom entrypoint to inject env vars
ENTRYPOINT ["/entrypoint.sh"]