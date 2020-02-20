USE POCDB;


# Dump of table USER_TABLE
# ------------------------------------------------------------

CREATE TABLE IF NOT EXISTS `USER_TABLE` (
  `name` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `public_key` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
