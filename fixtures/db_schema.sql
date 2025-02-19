-- booksテーブルを追加
CREATE TABLE `books` (
`id` BIGINT (10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
`title` VARCHAR (255) NOT NULL COMMENT 'タイトル',
`price` BIGINT (10) NOT NULL COMMENT '価格',
`author` VARCHAR (255) NOT NULL COMMENT '著者',
`book_img_url` TEXT NOT NULL COMMENT '本URL',
`rakuten_url` TEXT NOT NULL COMMENT '楽天ページへのURL',
`rakuten_review` INT NOT NULL COMMENT '楽天での評価',
`isbn` BIGINT (10) NOT NULL COMMENT '楽天APIから発行される一意なID',
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日',
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
PRIMARY KEY (`id`),
UNIQUE (`isbn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- usersテーブルを追加
CREATE TABLE `users` (
`id` BIGINT (10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
`email` VARCHAR (255) NOT NULL COMMENT 'メールアドレス',
`display_name` VARCHAR (255) NOT NULL COMMENT '表示名',
`avatar_url` TEXT NOT NULL COMMENT 'プロフィール画像',
`about` TEXT DEFAULT NULL COMMENT '自由記述欄',
`total_price` BIGINT DEFAULT 0 COMMENT '読んだ本の合計金額',
`recommendation_book_id` BIGINT (10) DEFAULT 1,
`is_terms_of_service` TINYINT (1) DEFAULT '0' COMMENT 'ユーザが利用規約に同意したか否かを表示',
`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日',
`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
PRIMARY KEY (`id`),
UNIQUE (`email`)
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


INSERT INTO `books` (`id`, `title`, `price`, `author`, `book_img_url`, `rakuten_url`, `rakuten_review`, `isbn`)
VALUES
(1, '超集中力', 1000, 'hoge', 'hogehoge@@hogehoge.com', '楽天のURL.com', 3, 23);

INSERT INTO `users` (`id`, `email`, `display_name`, `avatar_url`, `about`, `total_price`, `recommendation_book_id`, `is_terms_of_service`)
VALUES
(1, 'yata62885@gmail.com', 'tatsuya', 'avatar_url.img', '僕はyamamura', 0, 1, 1);

INSERT INTO `user_identities` (`id`, `sub`, `user_id`)
VALUES
(1, 'qwertsddffg', 1);

INSERT INTO `users_books_registrations` (`id`, `is_action_completed`, `user_id`, `book_id`)
VALUES
(1, 0, 1, 1);

INSERT INTO `actions` (`id`, `is_finished`, `user_book_registration_id`, `content`)
VALUES
(1, 0, 1, '早起きする'),
(2, 0, 1, 'ご飯食べる'),
(3, 0, 1, 'ランニングする');
