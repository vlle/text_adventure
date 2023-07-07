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


INSERT INTO image (name, emoji) VALUES
('bug', '🐛'),
('dragon', '🐉'),
('ghost', '👻'),
('alien', '👽'),
('robot', '🤖'),
('monster', '👾'),
('key', '🔑'),
('door', '🚪'),
('bed', '🛏️'),
('detective', '🕵️‍♂️');


CREATE TABLE IF NOT EXISTS location (
  id          serial primary key,
  title       varchar(50) not null,
  image_id    integer references image(id),
  description varchar(50) not null,
  xy          point,
  constraint unique_title unique (title)
);

CREATE TABLE IF NOT EXISTS item (
  id         serial primary key,
  name       varchar(50) not null,
  description varchar(50) not null,
  location_id integer references location(id),
  image_id    integer references image(id),
  CONSTRAINT unique_name UNIQUE (name)
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

