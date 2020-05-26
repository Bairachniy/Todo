begin;

CREATE TABLE if not exists todos_tmp (
    id    varchar(255) primary key default uuid_generate_v4(),
    todo  text NOT NULL UNIQUE
);

insert into todos_tmp (id, todo)
select cast(id as varchar(255)), todo from todos;

drop table todos;

alter table todos_tmp rename to todos;

commit;