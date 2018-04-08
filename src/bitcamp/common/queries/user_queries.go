package queries

import(
	sq "github.com/Masterminds/squirrel"
	"log"
	"errors"
	"bitcamp/common/models"
	"fmt"
)

func CreateUser(name string, isParent bool) error {
	query := sq.Insert("user").Columns("name", "isParent").
		Values(name, isParent)

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

func UpdateUserFCM(username, fcmId string) error{
	query := sq.Update("user").Set("fcm_id", fcmId).
		Where(sq.Eq{"name": username})
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

func GetUserById(id int) (*models.User, error) {
	query := sq.Select("*").From("user").Where(sq.Eq{"id": id})
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

	userScan, err  := models.ScanUsers(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(userScan) > 0{
		return &userScan[0], nil
	}else{
		return nil, errors.New("not found")
	}
}

func GetUserByName(name string) (*models.User, error){
	query := sq.Select("*").From("user").Where(sq.Eq{"name": name})
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

	userScan, err  := models.ScanUsers(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(userScan) > 0{
		return &userScan[0], nil
	}else{
		return nil, errors.New("not found")
	}
}

func PairParentAndChild(p_id, c_id int) (error){
	query := sq.Insert("parent_child_relation_table").Columns("c_id", "p_id").
		Values(p_id, c_id)

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

func GetParentsOfChildByName(id int) ([]models.User, error){

	query := fmt.Sprintf("select id, name, isParent, fcm_id from user left outer join parent_child_relation_table pcrt on user.id = pcrt.p_id where pcrt.c_id=%v;", id)

	log.Println(query)

	db, err := GetDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows , err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	log.Println("this is where it fails i bet")

	scans, err  := models.ScanUsers(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return scans, nil
}
