CREATE TABLE `blog_tag` (
    `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `created_at` int unsigned NOT NULL,
    `created_by` varchar(100) NOT NULL,
    `updated_at` int unsigned NOT NULL,
    `updated_by` varchar(100) NOT NULL,
    `deleted_at` int unsigned,
    `name` varchar(100) DEFAULT '' NOT NULL,
    `state` tinyint unsigned DEFAULT '1' NOT NULL COMMENT '0: disable, 1: enable'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `blog_article` (
    `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `created_at` int unsigned NOT NULL,
    `created_by` varchar(100) NOT NULL,
    `updated_at` int unsigned NOT NULL,
    `updated_by` varchar(100) NOT NULL,
    `deleted_at` int unsigned,
    `title` varchar(100) DEFAULT '' NOT NULL,
    `desc` varchar(255) DEFAULT '',
    `content` longtext,
    `cover_img_url` mediumtext,
    `state` tinyint DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `blog_article_tag` (
    `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `created_at` int unsigned NOT NULL,
    `created_by` varchar(100) NOT NULL,
    `updated_at` int unsigned NOT NULL,
    `updated_by` varchar(100) NOT NULL,
    `deleted_at` int unsigned,
    `article_id` int unsigned NOT NULL,
    `tag_id` int unsigned NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `blog_article_tag` ADD CONSTRAINT fk_article_id FOREIGN KEY (article_id) REFERENCES blog_article(id);
ALTER TABLE `blog_article_tag` ADD CONSTRAINT fk_tag_id FOREIGN KEY (tag_id) REFERENCES blog_tag(id);