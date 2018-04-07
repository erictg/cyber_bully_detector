package queries

import(
	sq "github.com/Masterminds/squirrel"
	"log"
	"errors"
	"bitcamp/common/models"
)

func InsertFlaggedText(content string, confidence float32, sent bool, userId int) error{
	query := sq.Insert("flagged_text").Columns("content", "confidence", "sent", "userId").
		Values(content, confidence, sent, userId)
	db, err := GetDB()
	if err != nil {
		log.Println(err)
		return err
	}

	result , err := query.RunWith(db).Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	if rows, err := result.RowsAffected(); err != nil || rows == 0{
		log.Println(err)
		log.Println(rows)
		return errors.New("failed to insert")
	}

	return nil
}

func GetTextById(id int) (*models.FlaggedText, error){
	query := sq.Select("*").From("flagged_text").Where(sq.Eq{"id": id})

	db, err := GetDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows , err := query.RunWith(db).Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	log.Println("this is where it fails i bet")

	scans, err  := models.ScanFlaggedTexts(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(scans) > 0{
		return &scans[0], nil
	}else{
		return nil, errors.New("not found")
	}
}

func GetTextsFromUser(id int) ([]models.FlaggedText, error){
	query := sq.Select("*").From("flagged_text").Where(sq.Eq{"user_id": id})

	db, err := GetDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows , err := query.RunWith(db).Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	log.Println("this is where it fails i bet")

	scans, err  := models.ScanFlaggedTexts(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return scans, nil
}