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

create table product_category
(
    id bigint auto_increment,
    parent_id bigint default 0 null comment '上级分类ID，0表示一级分类',
    name varchar(64) null comment '分类名称',
    level int default 0 null comment '分类级别： 0->1级 1->2级 ',
    show_status int default 0 null comment '是否展示 0:不展示，1:展示',
    sort int null comment '展示排序',
    keywords varchar(255) null comment '关键字',
    description text null comment '分类描述',
    constraint product_category_pk
        primary key (id)
)
    comment '产品分类';

alter table product_category
    add leaf_node int default 0 null comment '是否为叶子节点分类，0:是，1:否';

create table product
(
    id bigint auto_increment,
    brand_id bigint null comment '品牌ID',
    brand_name varchar(255) null comment '品牌名称',
    product_category_id bigint not null comment '产品分类ID',
    product_category_name varchar(255) null comment '商品分类名称',
    name varchar(64) null comment '商品名称',
    product_sn varchar(64) null comment '商品编号',
    delete_status int default 0 null comment '是否软删除，0:未删除，1:已删除',
    publish_status int default 0 null comment '上架状态，0:下架，1:上架',
    sort int null comment '商品展示排序',
    sales int default 0 null comment '销量',
    price decimal(10,2) null comment '商品价格',
    original_price decimal(10,2) null comment '市场价格/展示价格',
    stock int default 0 null comment '库存',
    detail_html text null comment '商品详情HTML',
    detail_mobile_html text null comment '商品详情手机端HTML',
    create_time datetime null,
    update_time datetime null,
    constraint product_pk
        primary key (id)
)
    comment '商品';


create table brand
(
    id bigint auto_increment,
    name varchar(64) null comment '品牌名称',
    sort int null comment '展示排序',
    logo varchar(255) null comment 'logo图片',
    constraint brand_pk
        primary key (id)
)
    comment '品牌';

