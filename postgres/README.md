## Postgres Driver

go get github.com/lib/pq

## Running Postgres

* Use the docker image on dockerhub

<pre>
docker pull postgres
docker run --name eventstoredb -e POSTGRES_PASSWORD=password -p 5432:5432  postgres
</pre>

* Connect via the command line:

<pre>
docker run -it --rm --link eventstoredb:postgres postgres psql -h postgres -U postgres
</pre>

## Set Up for the Sample

<pre>

create user esuser with password 'password';
create database esdb;

\c esdb

create table sample (
    name varchar(128) primary key,
    value varchar(256)
);

grant select, update, insert, delete on sample to esuser;
</pre>

