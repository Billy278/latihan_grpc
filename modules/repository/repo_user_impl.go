package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"latihan_grpc/modules/model"
)

type RepoUserImpl struct {
	DB *sql.DB
}

func NewRepoUserImpl(db *sql.DB) RepoUser {
	return &RepoUserImpl{
		DB: db,
	}
}

func (repo *RepoUserImpl) Show(ctx context.Context) (userRes []model.User, err error) {
	fmt.Printf("possition  Show %T \n", repo)
	sql := "SELECT id,name,email,age,jenkel FROM dataUser"
	row, err := repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	user := model.User{}
	for row.Next() {
		err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Jenkel)
		if err != nil {
			return
		}
		userRes = append(userRes, user)
	}
	return
}
func (repo *RepoUserImpl) Create(ctx context.Context, userIn model.User) (userRes model.User, err error) {
	fmt.Printf("possition  Create %T \n", repo)
	sql := "INSERT INTO dataUser(name,email,age,jenkel) VALUES($1,$2,$3,$4)"
	_, err = repo.DB.ExecContext(ctx, sql, userIn.Name, userIn.Email, userIn.Age, userIn.Jenkel)
	if err != nil {
		return
	}
	userRes.Name = userIn.Name
	return
}
func (repo *RepoUserImpl) FindById(ctx context.Context, userId uint64) (userRes model.User, err error) {
	fmt.Printf("possition  FindById %T \n", repo)
	sql := "SELECT id,name,email,age,jenkel FROM dataUser WHERE id=$1"
	row, err := repo.DB.QueryContext(ctx, sql, userId)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&userRes.Id, &userRes.Name, &userRes.Email, &userRes.Age, &userRes.Jenkel)
		if err != nil {
			return
		}
		return
	} else {
		err = errors.New("USER NOT FOUND")
		return
	}

}
func (repo *RepoUserImpl) Update(ctx context.Context, userIn model.User) (userRes model.User, err error) {
	fmt.Printf("possition  Update %T \n", repo)
	sql := "SELECT id,name,email,age,jenkel FROM dataUser WHERE id=$1 FOR UPDATE"
	tx, err := repo.DB.Begin()
	if err != nil {
		return
	}

	row, err := tx.QueryContext(ctx, sql, userIn.Id)
	if err != nil {
		tx.Rollback()
		return
	}
	if !row.Next() {
		err = errors.New("USER NOT FOUND")
		tx.Rollback()
		return
	}
	row.Close()
	fmt.Println(userIn)
	sqlUpdate := "UPDATE dataUser SET name=$1,email=$2,age=$3,jenkel=$4 WHERE id=$5"
	_, err = tx.ExecContext(ctx, sqlUpdate, userIn.Name, userIn.Email, userIn.Age, userIn.Jenkel, userIn.Id)
	if err != nil {
		tx.Rollback()
		return
	}
	userRes.Name = userIn.Name
	tx.Commit()
	return

}
func (repo *RepoUserImpl) Delete(ctx context.Context, userId uint64) (err error) {
	fmt.Printf("possition  Update %T \n", repo)
	sql := "SELECT id,name,email,age,jenkel FROM dataUser WHERE id=$1"
	tx, err := repo.DB.Begin()
	if err != nil {
		return
	}
	row, err := tx.QueryContext(ctx, sql, userId)
	if err != nil {
		tx.Rollback()
		return
	}
	if !row.Next() {
		err = errors.New("USER NOT FOUND")
		tx.Rollback()
		return
	}

	sqlDelete := "DELETE FROM dataUser WHERE id=$1"
	_, err = repo.DB.ExecContext(ctx, sqlDelete, userId)
	if err != nil {
		return
	}
	tx.Commit()
	return
}

//kesimplannya meskipun kita menggunakan execcontext dalam mengeskekusi
//updatenya akan tetapi sebernarnya menghasilskan row, berbeda dengan delete
//oleh sebab itu kita harus close row hasil findbyidnya
