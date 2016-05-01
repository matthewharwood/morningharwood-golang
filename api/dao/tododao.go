package tododao

import (
  "errors"
  todo "github.com/matthewharwood/morningharwood/api/model"
  "github.com/matthewharwood/morningharwood/config"
  "gopkg.in/mgo.v2/bson"
  "fmt"
)

const col string = "todos"

func All() (todo.Todos, error) {
  db := dbconfig.DB{}
  ts := todo.Todos{}
  fmt.Println(ts)
  s, err := db.DoDial()

  if err != nil {
    return ts, errors.New("There was an error trying to connect with the DB.")
  }

  defer s.Close()

  c := s.DB(db.Name()).C(col)

  err = c.Find(bson.M{}).All(&ts)

  if err != nil {
    return ts, errors.New("There was an error trying to find the todos.")
  }

  return ts, err
}