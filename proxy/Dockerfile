FROM nginx

COPY etc/nginx/conf.d/* /etc/nginx/conf.d/

# RUN cat /etc/nginx/nginx.conf

ENTRYPOINT /usr/sbin/nginx -g 'daemon off;' -c /etc/nginx/conf.d/proxy.conf
