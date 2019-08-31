FROM nginx:1.13.8
EXPOSE 8000
ADD nginx.conf /etc/nginx/conf.d/default.conf
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh
CMD ["/wait-for-it.sh", "image_processing:3000", "--", "nginx", "-g", "daemon off;"]
