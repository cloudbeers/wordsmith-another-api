FROM alpine:edge

COPY dispatcher /
COPY static /static/
COPY VERSION /

EXPOSE 80
CMD ["/dispatcher"]