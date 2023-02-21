create table blog_article
(
    id              int unsigned auto_increment
        primary key,
    tag_id          int unsigned     default '0' null comment '标签ID',
    title           varchar(100)     default ''  null comment '文章标题',
    `desc`          varchar(255)     default ''  null comment '简述',
    content         text                         null comment '内容',
    cover_image_url varchar(255)     default ''  null comment '封面图片地址',
    created_on      timestamp                    not null comment '新建时间',
    created_by      varchar(100)     default ''  null comment '创建人',
    modified_on     timestamp                    not null comment '修改时间',
    modified_by     varchar(255)     default ''  null comment '修改人',
    deleted_on      timestamp                    null,
    state           tinyint unsigned default '1' null comment '删除时间'
)
    comment '文章管理';

create table blog_auth
(
    id       int unsigned auto_increment
        primary key,
    username varchar(50) default '' null comment '账号',
    password varchar(50) default '' null comment '密码'
);

create table blog_tag
(
    id          int unsigned auto_increment
        primary key,
    name        varchar(100)     default ''  null comment '标签名称',
    created_on  timestamp                    null comment '创建时间',
    created_by  varchar(100)     default ''  null comment '创建人',
    modified_on timestamp                    null comment '修改时间',
    modified_by varchar(100)     default ''  null comment '修改人',
    deleted_on  int unsigned     default '0' null comment '删除时间',
    state       tinyint unsigned default '1' null comment '状态 0为禁用、1为启用'
)
    comment '文章标签管理';

