package models


type Brainee struct {
	ID int `json: "id"`
	Author string `json: "author"`
	Text string `json: "text"`
	Brand string `json: "brand"`
}

// FindBrainee finds a brainee in the DB with the ID
func FindBrainee(id int) (Brainee, error) {
	var brainee Brainee
	err := db.QueryRow("SELECT * FROM brainees WHERE id=?", id).Scan(&brainee.ID, &brainee.Author, &brainee.Text, &brainee.Brand)
	if err != nil {
		return brainee, err
	}
	return brainee, nil
}

// PostBrainee posts a new brainee in the DB
func PostBrainee(author, text, brand string) (bool, error) {
	insertQuery, err := db.Query("INSERT INTO brainees(author, text, brand) VALUES(?, ?, ?)", author, text, brand)
	if err != nil {
		return false, err
	} 

	defer insertQuery.Close()
	return true, nil
}