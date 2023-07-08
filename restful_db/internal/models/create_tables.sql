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
  image_id   integer references image(id)
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
(11, 'detective', '🕵️‍♂️');


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
INSERT INTO location (title, image_id, description, xy) VALUES
('Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(0,0)),
('Door', 8, 'Big, old wooden door. Seems like it wont open without a key', point(0,1)),
('Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(0,2)),
('Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(1,0)),
('Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(1,1)),
('Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(1,2)),
('Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(2,0)),
('Bed', 9, 'Very worn-out bed with blood stains on it. But there is no other bed, so..', point(2,1)),
('Stone floor', 10, 'This is cold, dirty stone floor. You do not like it.', point(2,2));

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
  location_id integer references location(id),
  image_id    integer references image(id),
  CONSTRAINT unique_name UNIQUE (name)
);

INSERT INTO monster (id, name, description, location_id, image_id) VALUES
(1, 'Sleeping zombie',  'It is sleeping, do not wake it up', 1, 3);

CREATE TABLE IF NOT EXISTS living_entity (

)
