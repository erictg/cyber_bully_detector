drop DATABASE if EXISTS bully;
CREATE DATABASE IF NOT EXISTS bully;
use bully;

create table user(
  id int auto_increment primary key,
  name varchar(255) not null,
  isParent bool default false,
  fcm_id varchar(255) default ''
);


create table parent_child_relation_table(
  p_id int,
  c_id int,
  constraint foreign key (p_id) references user(id) on delete cascade ,
  constraint foreign key (c_id) references user(id) on delete cascade
);

create table flagged_text(
  id int auto_increment primary key ,
  content varchar(1024) default '',
  confidence float default -1,
  sent bool default false ,
  user_id int default -1,
  other_number varchar(50) default '',
  constraint foreign key (user_id) references user(id) on delete set null
);

insert into user (id, name, isParent) values
  (3, 'sarah', false),
  (1, 'max', true),
  (2, 'gretchen', true );

insert into parent_child_relation_table (p_id, c_id) values
  (1, 3),
  (2, 3);