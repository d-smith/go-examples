Simple example showing how to dynamically register additional
handlers based on requests to the program


<pre>
$ curl localhost:4000/foo
404 page not found
$ curl -X PUT localhost:4000/register?uri=foo
$ curl localhost:4000/foo
This is the foo handler
</pre>