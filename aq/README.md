
<pre>
create type pubagg as object (
aggregate_id varchar2(60),
version integer
);

begin
    dbms_aqadm.create_queue_table(
         queue_table=> 'pub_qtab',
         queue_payload_type=>'esdbo.PUBAGG');
end;

begin
    dbms_aqadm.create_queue(
         queue_name=>'pubagg',
         queue_table=>'pub_qtab'
    );
end;

begin
    dbms_aqadm.start_queue(queue_name=>'pubagg');
end;
</pre>


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

</pre>

