package daos

import (
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/srinath-intelops/test-project-srinath-1/user-svc-1/pkg/rest/server/daos/clients/sqls"
	"github.com/srinath-intelops/test-project-srinath-1/user-svc-1/pkg/rest/server/models"
)

type UserdetailsDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateUserdetails(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS userdetails(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Username TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewUserdetailsDao() (*UserdetailsDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateUserdetails(sqlClient)
	if err != nil {
		return nil, err
	}
	return &UserdetailsDao{
		sqlClient,
	}, nil
}

func (userdetailsDao *UserdetailsDao) CreateUserdetails(m *models.Userdetails) (*models.Userdetails, error) {
	insertQuery := "INSERT INTO userdetails(Username)values(?)"
	res, err := userdetailsDao.sqlClient.DB.Exec(insertQuery, m.Username)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("userdetails created")
	return m, nil
}

func (userdetailsDao *UserdetailsDao) UpdateUserdetails(id int64, m *models.Userdetails) (*models.Userdetails, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	userdetails, err := userdetailsDao.GetUserdetails(id)
	if err != nil {
		return nil, err
	}
	if userdetails == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE userdetails SET Username = ? WHERE Id = ?"
	res, err := userdetailsDao.sqlClient.DB.Exec(updateQuery, m.Username, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("userdetails updated")
	return m, nil
}

func (userdetailsDao *UserdetailsDao) DeleteUserdetails(id int64) error {
	deleteQuery := "DELETE FROM userdetails WHERE Id = ?"
	res, err := userdetailsDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("userdetails deleted")
	return nil
}

func (userdetailsDao *UserdetailsDao) ListUserdetails() ([]*models.Userdetails, error) {
	selectQuery := "SELECT * FROM userdetails"
	rows, err := userdetailsDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var userdetails []*models.Userdetails
	for rows.Next() {
		m := models.Userdetails{}
		if err = rows.Scan(&m.Id, &m.Username); err != nil {
			return nil, err
		}
		userdetails = append(userdetails, &m)
	}
	if userdetails == nil {
		userdetails = []*models.Userdetails{}
	}

	log.Debugf("userdetails listed")
	return userdetails, nil
}

func (userdetailsDao *UserdetailsDao) GetUserdetails(id int64) (*models.Userdetails, error) {
	selectQuery := "SELECT * FROM userdetails WHERE Id = ?"
	row := userdetailsDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Userdetails{}
	if err := row.Scan(&m.Id, &m.Username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("userdetails retrieved")
	return &m, nil
}
