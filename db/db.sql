drop DATABASE if EXISTS bully;
CREATE DATABASE IF NOT EXISTS bully;
use bully;

create table user(
  id int auto_increment primary key,
  name varchar(255) not null,
  isParent bool default false
);

create table parent_child_relation_table(
  p_id int,
  c_id int,
  constraint foreign key (p_id) references user(id) on delete cascade ,
  constraint foreign key (c_id) references user(id) on delete cascade
);

create table flagged_text(
  id int auto_increment primary key ,
  content varchar(1024),
  confidence float,
  sent bool,
  user_id int,
  constraint foreign key (user_id) references user(id) on delete set null
);