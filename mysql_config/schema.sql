USE POCDB;


# Dump of table USER_TABLE
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `USER_TABLE` (
  `name` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT ''
  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;


insert into POCDB.USER_TABLE values ('SUNIL','TEST1@GMAIL.COM')

