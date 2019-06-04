# production environment
FROM nginx:1.15.9

COPY frontend/build /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d
EXPOSE 8000
CMD ["nginx", "-g", "daemon off;"]