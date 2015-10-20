CREATE TABLE `users` (
  id int(11) NOT NULL,
  name varchar(255),
  score int(11) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB;

INSERT INTO users VALUES (1, "aaa", 100);
INSERT INTO users VALUES (2, "bbb", 70);
INSERT INTO users VALUES (3, "ccc", 50);