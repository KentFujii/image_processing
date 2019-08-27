FROM minio/minio
ENV MINIO_ACCESS_KEY image_processing
ENV MINIO_SECRET_KEY password
EXPOSE 9000
CMD ["minio", "server", "/data"]
COPY .keep /data/image_processing/
