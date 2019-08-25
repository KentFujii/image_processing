FROM minio/minio
ENV MINIO_ACCESS_KEY image_processing
ENV MINIO_SECRET_KEY password
EXPOSE 9000
RUN mkdir -p /data/sample
COPY sample**  /data/sample/
CMD ["minio", "server", "/data"]
