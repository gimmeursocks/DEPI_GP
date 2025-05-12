#!/bin/sh
envsubst < /var/www/localhost/htdocs/init.template.js > /var/www/localhost/htdocs/init.js
exec lighttpd -D -f /etc/lighttpd/lighttpd.conf
