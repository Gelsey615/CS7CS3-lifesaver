package user

import (
  "db"
  "encoding/json"
  "errors"
)

func RegisterNewUser(u *db.User) error {
  check, _ := db.QueryDB(db.ColUser, "{\"eq\": \""+ u.Email +"\", \"in\": [\"email\"]}")
  if check != "{}" {
    return errors.New("user with email "+u.Email+" already exists")
  }
  doc, err := json.Marshal(u)
  if err != nil {
    return err
  }
  _, err = db.InsertToDB(db.ColUser, string(doc))
  return err
}

func Login(u *db.User) error {
  doc, err := db.QueryDB(db.ColUser, "{\"eq\": \""+ u.Email +"\", \"in\": [\"email\"]}")
  if err != nil || doc == "{}" {
    return errors.New("no user with email "+u.Email)
  }
  var users map[string]db.User
  if err := json.Unmarshal([]byte(doc), &users); err != nil {
    return err
  }
  for _, user := range users {
    if user.Pwd == u.Pwd {
      return nil
    }
  }
  return errors.New("wrong user pwd")
}
