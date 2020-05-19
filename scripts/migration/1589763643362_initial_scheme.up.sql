create table IF NOT EXISTS recipe
(
  id bigserial not null primary key,
  title varchar(255) not null
);