###Setup - Mac

Note: you need to check the paths in the pkgconfig file you reference via PKG_CONFIG_PATH

Also note: after you compile the binary, the only environment variable needed is
DYLD_LIBRARY_PATH (or presumably LD_LIBRARY_PATH on Linux)

And another thing... looks like if your oracle config changes you need to update
the pkgconfig, then remove and reinstall go-oci8

* Grab the oracle instaclient, for example from (here)[http://www.oracle.com/technetwork/topics/intel-macsoft-096467.html] 
* Make sure to download the 64-bit version
* Install them in /usr/local
* In the instant client directory, create a symbolic link of libtclntsh, e.g. `sudo ln -s libclntsh.dylib.11.1 libclntsh.dylib`
* export PKG_CONFIG_PATH=$GOPATH/src/github.com/d-smith/go-examples/oracle/pkgconfig/
* go get github.com/rjeczalik/pkgconfig/cmd/pkg-config
* go get -v github.com/mattn/go-oci8
* Note before go get that the install location and paths in oci8.pc must be aligned.

To run on the mac set the DYLD_LIBRARY_PATH , e.g. `export DYLD_LIBRARY_PATH=/usr/local/instantclient_11_2`

###Setup - Ubuntu 14

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

