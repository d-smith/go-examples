This example assumes that Apache Ignite has been started with the rest-http library copied
from lib/optional to lib, and ignite.sh run with a cache-enable config, e.g.

<pre>
bin/ignite.sh -v examples/config/example-cache.xml
</pre>

Probably the easiest way to start ignite is via the docker image. To run
using the GridGain community edition image:

<pre>
docker pull gridgain/gridgain-com

docker run -it --net=host -v `pwd`:/ignite  -e "CONFIG_URI=file:///ignite/examples/config/example-cache.xml" -e "OPTION_LIBS=ignite-rest-http" -p 8080:8080 -e "IGNITE_QUIET=false"  gridgain/gridgain-com
</pre>
