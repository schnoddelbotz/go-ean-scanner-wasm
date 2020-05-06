FROM nginx

# RUN sed -ie 's#.*text/mathml.*#application/wasm wasm;#' /etc/nginx/mime.types

COPY web /usr/share/nginx/html
COPY ngninx.conf /etc/nginx/nginx.conf
