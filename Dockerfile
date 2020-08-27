FROM alpine
ADD grpcProject /grpcProject
ENTRYPOINT [ "/grpcProject" ]
