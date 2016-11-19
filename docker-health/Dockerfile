FROM ubuntu:trusty

# Copy apt.conf.template to apt.conf with your proxy config if needed,
#o therwise comment out the following.
COPY apt.conf /etc/apt/apt.conf
RUN apt-get update && apt-get install -y curl && apt-get clean && rm -rf /var/lib/apt/lists

COPY main /opt/
EXPOSE 4000
HEALTHCHECK CMD curl --fail http://localhost:4000/health || exit 1
CMD ["/opt/main"]

