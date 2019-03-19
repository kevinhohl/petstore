package database

import (
	"fmt"

	"github.com/kevinhohl/petstore/pkg/model"
)

func FindPetByStatus(status string) []model.Pet {
	db := GetDBConn()

	query := `
		WITH tags as (
			SELECT pet_id, pet_tag_id, pet_tag_name
			FROM pet_tags
			JOIN pet_tag USING (pet_tag_id)
		)
		SELECT 
			pet_id,
			pet_name,
			pet_category_id,
			pet_category_name,
			pet_status_name,
			STRING_AGG(pet_tag_id::text, ', '::text) as tag_ids,
			STRING_AGG(pet_tag_name::text, ', '::text) as tag_names
		FROM pet 
		JOIN pet_category USING(pet_category_id)
		JOIN pet_status USING(pet_status_id)
		LEFT JOIN tags USING(pet_id)
		WHERE pet_status_name = $1
		GROUP BY pet_id, pet_name, pet_category_id, pet_category_name, pet_status_id, pet_status_name`

	var pets []model.Pet
	rows, err := db.Query(query, status)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var pet model.Pet
		err = rows.Scan(
			&pet.ID,
			&pet.Name,
			&pet.Category.ID,
			&pet.Category.Name,
			&pet.Status,
			&pet.TagsIDsRaw,
			&pet.TagsRaw,
		)
		pets = append(pets, pet.UnRaw())
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return pets
}

func FindPetByID(id int) (model.Pet, error) {
	db := GetDBConn()

	query := `
		WITH tags as (
			SELECT pet_id, pet_tag_id, pet_tag_name
			FROM pet_tags
			JOIN pet_tag USING (pet_tag_id)
		)
		SELECT 
			pet_id,
			pet_name,
			pet_category_id,
			pet_category_name,
			pet_status_name,
			STRING_AGG(pet_tag_id::text, ', '::text) as tag_ids,
			STRING_AGG(pet_tag_name::text, ', '::text) as tag_names
		FROM pet 
		JOIN pet_category USING(pet_category_id)
		JOIN pet_status USING(pet_status_id)
		LEFT JOIN tags USING(pet_id)
		WHERE pet_id = $1
		GROUP BY pet_id, pet_name, pet_category_id, pet_category_name, pet_status_id, pet_status_name`

	var pet model.Pet
	err := db.QueryRow(query, id).Scan(
		&pet.ID,
		&pet.Name,
		&pet.Category.ID,
		&pet.Category.Name,
		&pet.Status,
		&pet.TagsIDsRaw,
		&pet.TagsRaw,
	)
	if err != nil {
		return model.Pet{}, err
	}
	return pet.UnRaw(), nil
}

func AddPet(pet model.Pet) error {
	db := GetDBConn()

	fmt.Println("inserting pet_category")
	query := `INSERT INTO pet_category (pet_category_id, pet_category_name)
		VALUES (DEFAULT, $1) RETURNING pet_category_id`

	row := db.QueryRow(query, pet.Category.Name)
	var categoryID int64
	err := row.Scan(&categoryID)
	if err != nil {
		return err
	}
	fmt.Printf("finish inserting pet_category : %d\n", categoryID)

	fmt.Println("inserting pet")
	query = `INSERT INTO pet (pet_id, pet_name, pet_category_id, pet_status_id)
		VALUES (DEFAULT, $1, $2, $3) RETURNING pet_id`
	row = db.QueryRow(query, pet.Name, categoryID, model.StatusToName[pet.Status])

	var petID int64
	err = row.Scan(&petID)
	if err != nil {
		return err
	}
	fmt.Printf("finish inserting pet : %d\n", petID)

	fmt.Println("inserting pet tags")
	for _, val := range pet.Tags {
		fmt.Println("inserting pet tag")

		query := `INSERT INTO pet_tag (pet_tag_id, pet_tag_name)
			VALUES (DEFAULT, $1) RETURNING pet_tag_id`
		row = db.QueryRow(query, val.Name)
		var tagID int64
		err = row.Scan(&tagID)
		if err != nil {
			return err
		}
		fmt.Printf("finish inserting pet_tag : %d\n", tagID)

		fmt.Println("inserting pet tags")
		query = `INSERT INTO pet_tags (pet_id, pet_tag_id)
			VALUES ($1, $2)`
		_, err := db.Exec(query, petID, tagID)
		if err != nil {
			return err
		}
		fmt.Printf("finish inserting pet tags: %d -> %d\n", petID, tagID)
	}

	fmt.Println("inserting pet photo")

	for _, val := range pet.PhotoUrls {
		query := `INSERT INTO pet_photo (pet_photo_id, pet_id, pet_photo_url)
			VALUES (DEFAULT, $1, $2)`
		_, err := db.Exec(query, petID, val)
		if err != nil {
			return err
		}
	}
	fmt.Println("finish inserting pet photo")

	return nil
}

func DeletePet(petID int) error {
	db := GetDBConn()

	query := `DELETE FROM pet WHERE pet_id = $1`
	_, err := db.Exec(query, petID)
	if err != nil {
		return err
	}
	return nil
}
