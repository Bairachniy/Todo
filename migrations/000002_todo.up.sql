begin;

ALTER TABLE todos
ADD CONSTRAINT todo UNIQUE (todo);

commit;