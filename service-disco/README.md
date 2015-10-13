The example assumes a vagrant image in which consul and mountebank can be run. See Vagrantfile for an example
(you'll need to set your proxy info as appropriate).

After starting vagrant, vagrant ssh to the box and start mountebank, e.g. `mb`

In another shell, run consul:

<pre>
docker run -p 8400:8400 -p 8500:8500 -p 8600:53/udp -h node1 progrium/consul -server -bootstrap -ui-dir /ui
</pre>

Once consul and mountebank are started, use the commands in [this gist](https://gist.github.com/d-smith/6654063d9550a3bd9241#file-consul-son-of-docker-image-txt) 
to set up the service endpoints, register services and health checks, etc.