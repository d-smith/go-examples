The premise was could we write a function that could enqueue
something in Oracle AQ such that we could call it from
golang via the sql package.

While we can call functions, those functions cannot do any DML statements, which
the underlying call to the dbms_aq package does.

The next thing to explore is to see if we can just call a procedure via the sql
interface.


<pre>
create type pubagg as object (
agg_id_and_ver varchar2(60)
);

begin
    dbms_aqadm.create_queue_table(
         queue_table=> 'pubagg_qtab',
         queue_payload_type=>'esdbo.PUBAGG');
end;

begin
    dbms_aqadm.create_queue(
         queue_name=>'pubagg',
         queue_table=>'pubagg_qtab'
    );
end;

begin
    dbms_aqadm.start_queue(queue_name=>'pubagg');
end;

begin
SYS.DBMS_AQADM.STOP_QUEUE('pubagg');
end;

begin
SYS.DBMS_AQADM.DROP_QUEUE('pubagg');
end;

begin
SYS.DBMS_AQADM.DROP_QUEUE_TABLE('pub_qtab');
end;

drop type pubagg

</pre>



Enqueue

<pre>
declare
   message pubagg;
   enqueue_options     DBMS_AQ.enqueue_options_t;
   message_properties  DBMS_AQ.message_properties_t;
   message_handle      RAW(16);
begin
  message := pubagg('123',1);
  DBMS_AQ.ENQUEUE(
      queue_name              => 'pubagg',
      enqueue_options         => enqueue_options,
      message_properties      => message_properties,
      payload => message,
      msgid                   => message_handle);
end;
</pre>

Dequeue

<pre>
SET SERVEROUTPUT ON
DECLARE
dequeue_options     DBMS_AQ.dequeue_options_t;
message_properties  DBMS_AQ.message_properties_t;
message_handle      RAW(16);
message             pubagg;
BEGIN
   dequeue_options.navigation := DBMS_AQ.FIRST_MESSAGE;
   DBMS_AQ.DEQUEUE(
      queue_name          =>     'pubagg',
      dequeue_options     =>     dequeue_options,
      message_properties  =>     message_properties,
      payload             =>     message,
      msgid               =>     message_handle);
   DBMS_OUTPUT.PUT_LINE('aggregate_id'|| message.aggregate_id);
   DBMS_OUTPUT.PUT_LINE('version: '||message.version);
   COMMIT;
END;
</pre>

Go callable functions
<pre>
---Input should be <aggregate id>:<version> e.g. asdasd-dasd-asdsd:22
create or replace function enqueue_aggversion(aggAndVer varchar2) return int is 
begin
  declare
     message pubagg;
     enqueue_options     DBMS_AQ.enqueue_options_t;
     message_properties  DBMS_AQ.message_properties_t;
     message_handle      RAW(16);
  begin
    message := pubagg(aggAndVer);
    DBMS_AQ.ENQUEUE(
        queue_name              => 'pubagg',
        enqueue_options         => enqueue_options,
        message_properties      => message_properties,
        payload => message,
        msgid                   => message_handle);
  end;
end;
</pre>

<pre>
CREATE OR REPLACE PROCEDURE test_proc (
  outval OUT VARCHAR2
)
AS
BEGIN
    outval := 'hello';
END;

SET SERVEROUTPUT ON
declare
  testout varchar2(60);
begin
TEST_PROC(testout);
DBMS_OUTPUT.PUT_LINE('testout: '||testout);
end;

CREATE OR REPLACE function test_function return varchar2 is
begin
    return 'yo';
end;

select esdbo.test_function() from dual


SET SERVEROUTPUT ON
create or replace function tf3(s varchar2) return int is
begin
  DBMS_OUTPUT.PUT_LINE(s);
  return 0;
end;

CREATE OR REPLACE PROCEDURE kaboom(
    aggSpec in varchar2
)
AS
BEGIN
  if aggSpec != 'foo' then
    raise_application_error(-20345, 'kaboom');
  end if;
    
END;

CREATE OR REPLACE PROCEDURE ENQUEUE_AGG_SPEC(
    aggSpec in varchar2
)
AS

declare
   message pubagg;
   enqueue_options     DBMS_AQ.enqueue_options_t;
   message_properties  DBMS_AQ.message_properties_t;
   message_handle      RAW(16);
begin
  message := pubagg(aggSpec);
  DBMS_AQ.ENQUEUE(
      queue_name              => 'pubagg',
      enqueue_options         => enqueue_options,
      message_properties      => message_properties,
      payload => message,
      msgid                   => message_handle);
end;

CREATE OR REPLACE FUNCTION DEQUEUE_AGG_SPEC return varchar2 is
begin
DECLARE
dequeue_options     DBMS_AQ.dequeue_options_t;
message_properties  DBMS_AQ.message_properties_t;
message_handle      RAW(16);
message             pubagg;
BEGIN
   dequeue_options.navigation := DBMS_AQ.FIRST_MESSAGE;
   DBMS_AQ.DEQUEUE(
      queue_name          =>     'pubagg',
      dequeue_options     =>     dequeue_options,
      message_properties  =>     message_properties,
      payload             =>     message,
      msgid               =>     message_handle);
    return pubagg.agg_id_and_ver;      
END;
end;

</pre>

