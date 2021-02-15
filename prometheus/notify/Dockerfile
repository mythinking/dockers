FROM alpine
WORKDIR /opt/go
ADD ./src/ src/
EXPOSE 19000
EXPOSE 19001
EXPOSE 19002
EXPOSE 19003
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
