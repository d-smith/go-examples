For the redis samples, install the Redigo client library, e.g.

`go get github.com/garyburd/redigo/redis`

Run Redis using Docker for quick access...

<pre>
docker pull redis
docker run -p 6379:6379 --name redis1 redis
</pre>