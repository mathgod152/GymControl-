package database

import (
	"GusLem/gymControll/internal/entity"
	"database/sql"

	"github.com/google/uuid"
)

type ExerciceRepository struct {
	Db *sql.DB
}

func NewExerciceRepository(db *sql.DB) *ExerciceRepository {
	return &ExerciceRepository{Db: db}
}

func (a *ExerciceRepository) Create(Exercice *entity.ExerciceEntity) (*entity.ExerciceEntity, error) {
	id := uuid.New().String()
	stmt, err := a.Db.Prepare("INSERT INTO Exercice (id, name, email, birthday, gender, acount_type) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(id, Exercice.Name, Exercice.Desc, Exercice.CreateUser, Exercice.ExerciceName)
	if err != nil {
		return nil, err
	}
	return &entity.ExerciceEntity{
		ID:           id,
		Name:         Exercice.Name,
		Desc:         Exercice.Desc,
		CreateUser:   Exercice.CreateUser,
		ExerciceName: Exercice.ExerciceName,
	}, nil
}

func (a *ExerciceRepository) FindAll(page, limit int, sort string) ([]*entity.ExerciceEntity, error) {
	rows, err := a.Db.Query("SELECT id, name, email, birthday, gender, acount_type FROM Exercice")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Exercices := []*entity.ExerciceEntity{}
	for rows.Next() {
		var id, name, desc, createUser, exerciceName string
		if err := rows.Scan(&id, &name, &desc, &createUser, &exerciceName); err != nil {
			return nil, err
		}
		Exercice := entity.ExerciceEntity{
			ID:           id,
			Name:         name,
			Desc:         desc,
			CreateUser:   createUser,
			ExerciceName: exerciceName,
		}
		Exercices = append(Exercices, &Exercice)
	}
	return Exercices, nil
}

func (a *ExerciceRepository) FindByID(id string) (*entity.ExerciceEntity, error) {
	rows, err := a.Db.Query("SELECT id, name, email, birthday, gender, acount_type FROM user WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Exercice := entity.ExerciceEntity{}
	for rows.Next() {
		var id, name, desc, createUser, exerciceName string
		if err := rows.Scan(&id, &name, &desc, &createUser, &exerciceName); err != nil {
			return nil, err
		}
		Exercice = entity.ExerciceEntity{
			ID:           id,
			Name:         name,
			Desc:         desc,
			CreateUser:   createUser,
			ExerciceName: exerciceName,
		}
	}
	return &Exercice, nil
}
