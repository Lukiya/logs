	CREATE TABLE `Clients` (`ID` varchar(100) NOT NULL,`DBPolicy` int(11) NOT NULL,`Level` int(11) NOT NULL,PRIMARY KEY (`ID`)) DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMPRESSION = 'zstd_1.3.8' REPLICA_NUM = 3 BLOCK_SIZE = 16384 USE_BLOOM_FILTER = FALSE TABLET_SIZE = 134217728 PCTFREE = 0;
	INSERT INTO `Clients` VALUES('DL',0,1);
	INSERT INTO `Clients` VALUES('OLX',1,1);