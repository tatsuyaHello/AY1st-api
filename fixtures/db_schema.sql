-- usersテーブルを追加
CREATE TABLE `users` (
`id` BIGINT (10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
`email` VARCHAR (255) NOT NULL COMMENT 'メールアドレス',
`display_name` VARCHAR (255) NOT NULL COMMENT '表示名',
`avartar_url` TEXT NOT NULL COMMENT 'プロフィール画像',
`about` TEXT DEFAULT NULL COMMENT '自由記述欄',
`recommendation_book` TEXT NULL COMMENT '推薦本URL',
`is_terms_of_service` TINYINT (1) DEFAULT '0' COMMENT 'ユーザが利用規約に同意したか否かを表示',
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日',
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
PRIMARY KEY (`id`),
UNIQUE (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- booksテーブルを追加
CREATE TABLE `books` (
`id` BIGINT (10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
`rakuten_id` VARCHAR (255) NOT NULL COMMENT '楽天APIから発行される一意なID',
`title` VARCHAR (255) NOT NULL COMMENT 'タイトル',
`price` BIGINT (10) NOT NULL COMMENT '価格',
`author` VARCHAR (255) NOT NULL COMMENT '著者',
`book_img_url` TEXT NOT NULL COMMENT '本URL',
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日',
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
PRIMARY KEY (`id`),
UNIQUE (`rakuten_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- users_books_registrationsテーブルを追加
CREATE TABLE `users_books_registrations` (
`id` BIGINT (10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
`is_action_completed` INT (1) DEFAULT '0' COMMENT '終了したか否か',
`user_id` BIGINT (10),
`book_id` BIGINT (10),
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日',
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
PRIMARY KEY (`id`),
CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
CONSTRAINT fk_book_id FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- actionsテーブルを追加
CREATE TABLE `actions` (
`id` BIGINT (10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
`is_finished` INT (1) DEFAULT '0' COMMENT '終了したか否か',
`user_book_registration_id` BIGINT (10),
`content` VARCHAR (255) NOT NULL COMMENT '内容',
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日',
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
PRIMARY KEY (`id`),
CONSTRAINT fk_user_book_registration_id FOREIGN KEY (user_book_registration_id) REFERENCES users_books_registrations (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- user_identitiesテーブルを追加
CREATE TABLE `user_identities` (
`id` BIGINT (10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
`sub` VARCHAR (255) NOT NULL COMMENT 'Cognitoからの値',
`user_id` BIGINT (10),
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日',
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
PRIMARY KEY (`id`),
CONSTRAINT fk_useri_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `books` (`id`, `rakuten_id`, `title`, `price`, `author`, `book_img_url`)
VALUES
(1, '12345', '超集中力', 1000, 'hoge', 'hogehoge@@hogehoge.com');

INSERT INTO `users` (`id`, `email`, `display_name`, `avartar_url`, `about`, `recommendation_book`, `is_terms_of_service`)
VALUES
(1, 'yata62885@gmail.com', 'tatsuya', 'avartar_url.img', '僕はyamamura', 'hogehoge@@hogehoge.com', 1);

INSERT INTO `user_identities` (`id`, `sub`, `user_id`)
VALUES
(1, 'qwertsddffg', 1);

INSERT INTO `users_books_registrations` (`id`, `is_action_completed`, `user_id`, `book_id`)
VALUES
(1, 0, 1, 1);

INSERT INTO `actions` (`id`, `is_finished`, `user_book_registration_id`, `content`)
VALUES
(1, 0, 1, '早起きする');
