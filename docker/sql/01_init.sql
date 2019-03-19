
CREATE TABLE IF NOT EXISTS pet_status
(
  pet_status_id   INTEGER PRIMARY KEY NOT NULL,
  pet_status_name TEXT    NOT NULL
);

INSERT INTO pet_status (pet_status_id, pet_status_name)
VALUES
 (10, 'available'),
 (20, 'pending'),
 (30, 'sold');

CREATE TABLE IF NOT EXISTS pet_category
(
  pet_category_id   SERIAL PRIMARY KEY NOT NULL,
  pet_category_name TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS pet_tag
(
  pet_tag_id   SERIAL PRIMARY KEY NOT NULL,
  pet_tag_name TEXT   NOT NULL
);

CREATE TABLE IF NOT EXISTS pet
(
  pet_id          SERIAL  PRIMARY KEY NOT NULL,
  pet_name        TEXT    NOT NULL,
  pet_category_id INTEGER REFERENCES pet_category(pet_category_id),
  pet_status_id   INTEGER REFERENCES pet_status(pet_status_id)
);

CREATE TABLE IF NOT EXISTS pet_tags
(
  pet_id     INTEGER REFERENCES pet(pet_id) ON DELETE CASCADE,
  pet_tag_id INTEGER REFERENCES pet_tag(pet_tag_id)
);

CREATE TABLE IF NOT EXISTS pet_photo
(
  pet_photo_id  SERIAL  PRIMARY KEY NOT NULL,
  pet_id        INTEGER REFERENCES pet(pet_id) ON DELETE CASCADE,
  pet_photo_url TEXT    NOT NULL
);
