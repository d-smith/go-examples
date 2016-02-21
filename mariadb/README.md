## Basic Setup

This code assumes the following setup.

<pre>
show databases;

create database sample;

create user sampleusr identified by 'samplepw';

create table if not exists sample.sample (
	name varchar(128) primary key,
    value varchar(256)
);

grant select, update, insert, delete
on sample.sample
to sampleusr;
</pre>