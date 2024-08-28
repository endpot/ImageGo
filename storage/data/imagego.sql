SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Image Unique Code',
  `delete_code` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Image Delete Code',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Image File Name',
  `ext` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Image Extension',
  `width` int NOT NULL COMMENT 'Image Width in Pixels',
  `height` int NOT NULL COMMENT 'Image Height in Pixels',
  `nsfw` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Normal or NSFW',
  `uploader_ip` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Image Uploader IP',
  `fingerprint` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Image Fingerprint',
  `save_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Image Save Name',
  `size` bigint NOT NULL COMMENT 'Image Size in Bits',
  `views` bigint unsigned NOT NULL COMMENT 'Image View Counts',
  `created_at` datetime DEFAULT NULL COMMENT 'Create Time',
  `updated_at` datetime DEFAULT NULL COMMENT 'Update Time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_code` (`code`),
  UNIQUE KEY `uni_delete_code` (`delete_code`),
  KEY `idx_save_name` (`save_name`),
  KEY `idx_fingerprint` (`fingerprint`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

SET FOREIGN_KEY_CHECKS = 1;
