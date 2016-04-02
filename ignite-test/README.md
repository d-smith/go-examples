## Ignite Test

A simple program to throw some load at ignite to simulate anticipated access patterns.

### Amazon Setup - Cluster Nodes

Install docker and pull the grid gain image.

<pre>
wget -qO- https://get.docker.com/ | sh
sudo usermod -aG docker ubuntu
docker pull gridgain/gridgain-com
</pre>

Next pull in the config and edit the particulars.

<pre>
wget https://raw.githubusercontent.com/d-smith/go-examples/master/ignite-test/aws-cache-config.xml
</pre>

Note in the above, the security groups need to permit traffic on ports 47100 and 47500 amongst the
cluster members.

To run the cache:

<pre>
docker run -it --net=host -v /home/ubuntu:/ignite -e "CONFIG_URI=file:///ignite/aws-cache-config.xml" -e "OPTION_LIBS=ignite-rest-http"  -p 8080:8080 -e "IGNITE_QUIET=false" gridgain/gridgain-com
</pre>