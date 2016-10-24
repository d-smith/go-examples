FROM scratch
COPY ca.crt /opt/
COPY service.crt /opt/
COPY service.key /opt/
COPY server /opt/
EXPOSE 51000
ENTRYPOINT ["/opt/server"]
CMD ["/opt/server","/opt/service.key","/opt/service.crt","/opt/ca.crt"]