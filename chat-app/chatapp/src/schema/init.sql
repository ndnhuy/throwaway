DROP TABLE IF EXISTS `message`;
CREATE TABLE message
(
    id                bigint(20) NOT NULL AUTO_INCREMENT,
    from_user             varchar(64) NOT NULL,
    to_user             varchar(64) NOT NULL,
    content             varchar(64) NOT NULL,
    created_at          datetime NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `message_seq`;
CREATE TABLE message_seq(ID INT NOT NULL);
INSERT INTO message_seq VALUES (0);