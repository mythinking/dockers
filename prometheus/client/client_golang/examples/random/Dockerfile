FROM alpine
WORKDIR /random
ADD ./random random
EXPOSE 18080
EXPOSE 18081
EXPOSE 18082
EXPOSE 18083
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
CMD ./random -listen-address=:18080
