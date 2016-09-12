## Working with Oracle using Golang

To work with Oracle using the `https://github.com/mattn/go-oci8` or
`https://github.com/rana/ora` drivers, you must:

1. Install the Oracle Instant Client package
2. Install the golang aware version of pkgconfig from https://godoc.org/github.com/rjeczalik/pkgconfig
3. Create a package configuration reflecting your Oracle set up
4. Install the Oracle driver
5. Optional - run Oracle locally using docker.

### Oracle Instant Client Setup - Mac

* Grab the oracle instaclient, for example from (here)[http://www.oracle.com/technetwork/topics/intel-macsoft-096467.html] 
* Make sure to download the 64-bit version
* Install them where ever you want, I installed mine in /opt/oracle, which is the path the rest of these
instructions assume.
* In the instant client directory, create a symbolic link of libtclntsh, e.g. `sudo ln -s libclntsh.dylib.11.1 libclntsh.dylib`
* Edit your .bach_profile or preferred mechanism for setting up your environment to ensure the 
DYLD_LIBRARY_PATH environment variables includes your instant client directory, for example
`export DYLD_LIBRARY_PATH=/opt/oracle/instantclient_12_1`

### Pkgconfig

* go get github.com/rjeczalik/pkgconfig/cmd/pkg-config
* Determine where you want to keep pkgconfig files. For working with the Oracle
Instant Client, you will need a file named oci8.pc.
* I keep my oci8.pc file in $HOME/bin, so my golang environment set up includes
`export PKG_CONFIG_PATH=$HOME/bin`

You can use the pkgconfig/oci8.pc file as a template for your set up. Note the 
paths in oci8.pc must match your Oracle installation or your set up will fail.

### Oracle Driver Installation

The installation of the Oracle driver depends on the proper installation of Oracle 
Instant Client and the Pkgconfig. If the driven installation via go get fails because
of instant client or pkgconfig configuration problems, or if you update or change instant
client or pkgconfig, you will need to reinstall the driver.

To install the driver:

* go get -v github.com/mattn/go-oci8


### Oracle Docker Image

Note: downloading [this image](https://hub.docker.com/r/sath89/oracle-12c/) can be problematic as it is huge. There's an image layer
that is 2.67 GB!

When running this in a Vagrant environment, I needed a virtual machine definition of 2048 MBytes - I had start up failures
when running in it a 1GB machine.

<pre>
docker pull sath89/oracle-12c
docker run -d -p 8080:8080 -p 1521:1521 sath89/oracle-12c
</pre>

With database files that persist across runs:

<pre>
docker run -p 8080:8080 -p 1521:1521 -v /opt/oradata:/u01/app/oracle sath89/oracle-12c
</pre>

For good measure, when running Oracle in Docker, use docker stop <image id> to 
cleanly shutdown you Oracle instance. If you are only running Oracle in a dedicated
Vagrant machine, you can just do the following everytime you wan to shutdown.

<pre>
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
</pre>

### Note for Junos VPN Users

Junos Pulse updates your /etc/host file each time you connect, which
means with SQL*Net you will see ORA-21561: OID generation failed

To avoid this error, edit
/etc/hosts and add your hostname after localhost, e.g.

<pre>
127.0.0.1 localhost [local computer name]
</pre>

### Oracle Instant Client Setup - Ubuntu 14

Grab the Oracle instant client - I grabbed the following from (here)[http://www.oracle.com/technetwork/topics/linuxx86-64soft-092277.html]

<pre>
instantclient-basic-linux.x64-12.1.0.2.0.zip
instantclient-sdk-linux.x64-12.1.0.2.0.zip
instantclient-sqlplus-linux.x64-12.1.0.2.0.zip
</pre>


Unzip the archives

<pre>
sudo mkdir -p /opt/oracle/instantclient_12_1
cd /opt/oracle
sudo cp /vagrant/*.zip .
sudo apt-get install unzip
sudo unzip instantclient-basic-linux.x64-12.1.0.2.0.zip 
sudo unzip instantclient-sdk-linux.x64-12.1.0.2.0.zip
sudo unzip instantclient-sqlplus-linux.x64-12.1.0.2.0.zip
</pre>

Make the links as per the instructions on the download page and update paths, e.g.

<pre>
cd instantclient_12_1
sudo ln -s libclntsh.so.12.1 libclntsh.so
sudo ln -s libocci.so.12.1 libocci.so
export LD_LIBRARY_PATH=/opt/oracle/instantclient_12_1:$LD_LIBRARY_PATH
export PATH=/opt/oracle/instantclient_12_1:$PATH
</pre>

Needed the following supporting library, might have just needed it for running 
sqlplus:

<pre>
sudo apt-get install libaio1 libaio-dev
</pre>

Then...

<pre>
export PKG_CONFIG_PATH=$GOPATH/src/github.com/d-smith/go-examples/oracle/pkgconfig/
go get github.com/rjeczalik/pkgconfig/cmd/pkg-config
go get -v github.com/mattn/go-oci8
</pre>

Note before go get that the install location and paths in oci8.pc must be aligned.
