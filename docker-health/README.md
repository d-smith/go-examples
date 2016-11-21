This project shows how to incorporate a docker health check. For this
example I can do a get on the endpoint the container exposes, but
more involved health checks that examine integration dependencies, connection
status, and other things needed by the app to assess health could also be
included.

For troubleshooting the health check curl, you can use docker inspect, e.g.

<pre>
docker inspect --format "{{json .State.Health }}" focused_bohr
</pre>

If building the image behind an http proxy, copy apt.conf.template to apt.conf and
edit apt.conf to include your proxy configuration.

Note the use of `--noproxy localhost` in the curl options - this is needed
if your docker daemon is configured with an http proxy.


