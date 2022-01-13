create table member
(
    id bigint auto_increment,
    username varchar(64) not null comment '用户名（唯一标识）',
    password varchar(64) not null comment '密码',
    nickname varchar(64) null comment '昵称',
    phone varchar(11) null comment '手机号',
    status int default 0 null comment '账号状态：0-启用 1-禁用',
    create_time datetime default now() null,
    update_time datetime null,
    constraint member_pk
        primary key (id)
)
    comment '会员表';

