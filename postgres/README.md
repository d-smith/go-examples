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

create sequence sampleseq increment by 1

create table sample (
    recno integer default nextval('sampleseq'),
    name varchar(128) primary key,
    value varchar(256),
    dablob bytea
);

grant select, update, insert, delete on sample to esuser;
grant select, update on sampleseq to esuser;
</pre>

