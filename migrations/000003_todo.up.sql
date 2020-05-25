begin;

alter table todos alter column id set default uuid_generate_v4();

commit;