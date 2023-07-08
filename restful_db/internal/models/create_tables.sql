-- postgresql dialect
-- # # #
-- # # #
-- # # #

CREATE TABLE IF NOT EXISTS image (
  id         serial primary key,
  name       varchar(50) not null,
  url        varchar(50),
  emoji      varchar(50) not null,
  CONSTRAINT unique_emoji UNIQUE (emoji)
);

CREATE TABLE IF NOT EXISTS p_user (
  id         serial primary key,
  name       varchar(50) not null,
  password   varchar(50) not null,


  hp         integer not null,
  fight_power integer not null,
  armor      integer not null,
  image_id   integer references image(id)
);

CREATE TABLE IF NOT EXISTS user_item (
  user_id    integer references p_user(id),
  item_id    integer references item(id)
);


INSERT INTO image (id, name, emoji) VALUES
(1, 'bug', '🐛'),
(2, 'dragon', '🐉'),
(3, 'ghost', '👻'),
(4, 'alien', '👽'),
(5, 'robot', '🤖'),
(6, 'monster', '👾'),
(7, 'key', '🔑'), 
(8, 'door', '🚪'),
(9, 'bed', '🛏️'),
(10, 'stone_floor', '🪨'),
(11, 'detective', '🕵️‍♂️'),
(12, 'zombie', '🧟');


CREATE TABLE IF NOT EXISTS location (
  id          serial primary key,
  title       varchar(50) not null,
  image_id    integer references image(id),
  description varchar(127) not null,
  xy          point
  -- constraint unique_title unique (title)
);

-- # # #
-- # # #
-- # # #
INSERT INTO location (id, title, image_id, description, xy) VALUES
(1, 'Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(0,0)),
(2, 'Door', 8, 'Big, old wooden door. Seems like it wont open without a key', point(0,1)),
(3, 'Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(0,2)),
(4, 'Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(1,0)),
(5, 'Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(1,1)),
(6, 'Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(1,2)),
(7, 'Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(2,0)),
(8, 'Bed', 9, 'Very worn-out bed with blood stains on it. But there is no other bed, so..', point(2,1)),
(9, 'Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(2,2));

CREATE TABLE IF NOT EXISTS item (
  id         serial primary key,
  name       varchar(50) not null,
  description varchar(50) not null,
  location_id integer references location(id),
  image_id    integer references image(id)
);

INSERT INTO item (name, description) VALUES
('Key', 'Old rusty key, it may open something important'),
('Hammer', 'Be careful, you might hit someone with it'),
('Lantern', 'You really should not turn light off');

CREATE TABLE IF NOT EXISTS monster (
  id         serial primary key,
  name       varchar(50) not null,
  description varchar(50) not null,
  hp         integer not null,
  fight_power integer not null,
  armor      integer not null,
  location_id integer references location(id),
  image_id    integer references image(id)
);

INSERT INTO monster (id, name, hp, fight_power, armor, description, location_id, image_id) VALUES
(1, 'Sleeping zombie',  10, 3, 0, 'It is sleeping, do not wake it up', 1, 12),
(2, 'Sleeping zombie',  10, 3, 0, 'It is sleeping, do not wake it up', 1, 12),
(3, 'Sleeping zombie',  10, 3, 0, 'It is sleeping, do not wake it up', 1, 12);
