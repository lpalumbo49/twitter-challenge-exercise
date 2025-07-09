CREATE TABLE IF NOT EXISTS user
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(32) NOT NULL,
    surname     VARCHAR(32) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    username    VARCHAR(32) NOT NULL,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT  unique_user_email UNIQUE (email),
    CONSTRAINT  unique_user_username UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS tweet
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    user_id    INT          NOT NULL,
    text       VARCHAR(280) NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_tweet_user_id FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE IF NOT EXISTS follower
(
    user_id             INT      NOT NULL,
    followed_by_user_id INT      NOT NULL,
    created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, followed_by_user_id),
    CONSTRAINT fk_follower_user_id FOREIGN KEY (user_id) REFERENCES user (id),
    CONSTRAINT fk_follower_followed_by_user_id FOREIGN KEY (followed_by_user_id) REFERENCES user (id)
);
