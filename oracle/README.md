Setup

* Grab the oracle instaclient, for example from (here)[http://www.oracle.com/technetwork/topics/intel-macsoft-096467.html] 
* Make sure to download the 64-bit version
* Install them in /usr/local
* In the instant client directory, create a symbolic link of libtclntsh, e.g. `sudo ln -s libclntsh.dylib.11.1 libclntsh.dylib`
* export PKG_CONFIG_PATH=$GOPATH/src/github.com/d-smith/go-examples/oracle/pkgconfig/
* go get -v github.com/mattn/go-oci8
* Note before go get that the install location and paths in oci8.pc must be aligned.

To run on the mac set the DYLD_LIBRARY_PATH , e.g. `export DYLD_LIBRARY_PATH=/usr/local/instantclient_11_2`
