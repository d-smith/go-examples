For the redis samples, install the go-redis client library, e.g.

`go get https://github.com/go-redis/redis`

Run Redis using Docker for quick access...

<pre>
docker pull redis
docker run -p 6379:6379 --name redis1 redis
</pre>

And the CLI

<pre>
docker run -it --link redis1:redis --rm redis redis-cli -h redis -p 6379
</pre>