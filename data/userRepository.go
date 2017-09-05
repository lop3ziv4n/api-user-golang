package data

import (
	"github.com/lop3ziv4n/api-user-golang-mongo/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) Create(user *models.User) error {
	obj_id := bson.NewObjectId()
	user.Id = obj_id
	err := r.C.Insert(&user)
	return err
}

func (r *UserRepository) GetAll() []models.User {
	var users []models.User
	iter := r.C.Find(nil).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *UserRepository) GetById(id string) (user models.User, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&user)
	return
}

func (r *UserRepository) GetAllByName(name string) []models.User {
	var users []models.User
	iter := r.C.Find(bson.M{"name": name}).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *UserRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *UserRepository) Update(id string, user *models.User) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &user)
	return err
}
