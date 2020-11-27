FROM nginx:1.19.4-alpine

EXPOSE 80

RUN rm -r /usr/share/nginx/html/

COPY static-content /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
