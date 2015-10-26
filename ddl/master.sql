-- CREATE DATABASE game_master CHARACTER SET utf8;
-- GRANT ALL PRIVILEGES ON `game_master`.* TO 'game'@'localhost';

DROP TABLE IF EXISTS `user_shard`;
CREATE TABLE `user_shard` (
  id int(11) NOT NULL,
  shard_id int(11) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB;

BEGIN;
INSERT INTO user_shard VALUES (1, 1);
INSERT INTO user_shard VALUES (2, 2);
INSERT INTO user_shard VALUES (3, 1);
COMMIT;