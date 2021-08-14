CREATE TABLE `danjuan_fund` (
    `id` int NOT NULL AUTO_INCREMENT,
    `fund_name` varchar(64) DEFAULT '' COMMENT '基金名称',
    `fund_code` varchar(16) NOT NULL DEFAULT '' COMMENT '基金代码',
    `managers` varchar(32) NOT NULL DEFAULT '' COMMENT '管理人',
    `end_date` varchar(32) NOT NULL DEFAULT '' COMMENT '季报日期',
    `detail_json` text NOT NULL COMMENT '蛋卷基金详细信息 json',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_code` (`fund_code`, `fund_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;