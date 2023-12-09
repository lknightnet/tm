FROM ubuntu

WORKDIR /task-manager

COPY config/config.yaml config/
COPY task-manager /task-manager

EXPOSE 8001

CMD ["./task-manager"]