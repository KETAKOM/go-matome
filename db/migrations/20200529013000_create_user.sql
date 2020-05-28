
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE user
(
  id int NOT NULL auto_increment COMMENT 'ID',
  name varchar(128) NOT NULL COMMENT '名前',
  CONSTRAINT pk_user PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user;
