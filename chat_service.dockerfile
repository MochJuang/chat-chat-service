LABEL AUTHOR=MochJuang
LABEL VERSION=1.0

FROM alpine:latest

RUN mkdir /app

COPY ChatService  /app

CMD ["/app/ChatService"]