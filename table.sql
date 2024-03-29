create table tbl_app_info(
    app_id int auto_increment primary key,
    app_name varchar (1024) not null ,
    app_type varchar (64)   not null ,
    create_time timestamp default current_timestamp ,
    develop_path varchar (256)  not null
)ENGINE=InnoDB  default charset=utf8 auto_increment=1;

create table tbl_app_ip(
    app_id int,
    ip varchar (64),
    key app_id_ip_index (app_id,ip)
)ENGINE=InnoDB  default charset=utf8 auto_increment=1;

create table tbl_log_info(
    log_id int auto_increment primary key ,
    app_id int not null ,
    log_path varchar (64) not null ,
    create_time timestamp default current_timestamp ,
    status tinyint default 1
)ENGINE=InnoDB  default charset=utf8 auto_increment=1;