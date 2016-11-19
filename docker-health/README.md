docker inspect --format "{{json .State.Health }}" focused_bohr

Even when building outside a proxy situration I had to run with 
no_proxy=localhost for the health check to be invoked - probably based 
on the docker daemon setting.

docker run -e no_proxy=localhost -p 4000:4000 echo

