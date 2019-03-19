INSERT INTO pet_category (pet_category_id, pet_category_name)
VALUES
 (DEFAULT, 'dog'),
 (DEFAULT, 'cat'),
 (DEFAULT, 'other');

INSERT INTO pet (pet_name, pet_category_id, pet_status_id)
VALUES
 ('doggie', 1, 10),
 ('catto', 2, 20),
 ('drago', 3, 30);

INSERT INTO pet_tag (pet_tag_id, pet_tag_name)
VALUES
 (DEFAULT, 'good_boy'),
 (DEFAULT, 'scratches'),
 (DEFAULT, 'breathes_fire'),
 (DEFAULT, 'black'),
 (DEFAULT, 'red'),
 (DEFAULT, 'green');

INSERT INTO pet_tags (pet_id, pet_tag_id)
VALUES
 (1, 1),
 (1, 4),
 (2, 2),
 (2, 5),
 (3, 3),
 (3, 6);

INSERT INTO pet_photo (pet_photo_id, pet_id, pet_photo_url)
VALUES
 (DEFAULT, 1, 'doggie_pic_1'),
 (DEFAULT, 1, 'doggie_pic_2'),
 (DEFAULT, 2, 'catto_pic_1'),
 (DEFAULT, 2, 'catto_pic_2'),
 (DEFAULT, 3, 'drago_pic_1'),
 (DEFAULT, 3, 'drago_pic_2');
