## Basic Setup

This code assumes the following setup.

<pre>
show databases;

create database roll;

create user rolluser identified by 'rollpw';

create table if not exists roll.sample (
	name varchar(128) primary key,
    value varchar(256)
);

grant select, update, insert, delete
on roll.sample
to rolluser;
</pre>