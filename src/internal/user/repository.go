package user

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	entityUser "github.com/notAlyosha/quiz-go/internal/entity/user"
	"github.com/notAlyosha/quiz-go/pkg/dbcontext"
	uuid "github.com/satori/go.uuid"
)

func createUserInStorage(FID string, salt string, newUser entityUser.UserCreate) error {
	err := dbcontext.GetDB().Transactional(func(tx *dbx.Tx) error {

		// create USER
		_, err := tx.Insert("users", dbx.Params{
			"front_id": FID,
			"role":     newUser.Role,
			"logo_URL": newUser.LogoURL,
		}).Execute()

		if err != nil {
			return err
		}

		// create user's PROFILE
		_, err = tx.Insert("profiles", dbx.Params{
			"login":   newUser.Login,
			"salt":    salt,
			"name":    newUser.Name,
			"email":   newUser.Email,
			"phone":   newUser.Phone,
			"logoURL": newUser.LogoURL,
		}).Execute()

		if err != nil {
			return err
		}

		// get USER ID to create user'spassword
		user := entityUser.User{}
		q := tx.NewQuery("SELECT id FROM users WHERE front_id={:id}")
		q.Bind(dbx.Params{"id": FID})

		err = q.One(&user)
		if err != nil {
			return err
		}

		// create user's PASSWORD in HISTORY
		_, err = tx.Insert("users-passwords-history", dbx.Params{
			"user_id": user.ID,
			"value":   newUser.Password,
		}).Execute()

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func updateUserInStorage(FID string, newUser entityUser.UserCreate) error {
	err := dbcontext.GetDB().Transactional(func(tx *dbx.Tx) error {
		var err error

		// update USER
		_, err = tx.Update("users", dbx.Params{
			"front_id": FID,
			"role":     newUser.Role,
			"logo_URL": newUser.LogoURL,
		},
			dbx.HashExp{"id": FID},
		).Execute()

		if err != nil {
			return err
		}

		// update user's PROFILE
		_, err = tx.Update("profiles", dbx.Params{
			"name":  newUser.Name,
			"email": newUser.Email,
			"phone": newUser.Phone,
		},
			dbx.HashExp{
				"user_id": FID,
			}).Execute()

		if err != nil {
			return err
		}

		// update user's PASSWORDS HISTORY

		q := dbcontext.GetDB().NewQuery("SELECT value FROM users-passwords-history WHERE user_id={:user_id}")
		q.Bind(dbx.Params{
			"user_id": newUser.Login,
		})

		passwordStruct := entityUser.UserPasswordHistroy{}
		err = q.One(&passwordStruct)
		if err != nil {
			return err
		}

		if passwordStruct.Value == newUser.Password {
			return nil
		}

		_, err = tx.Update("users-passwords-history", dbx.Params{
			"user_id": newUser.Name,
			"value":   newUser.Password,
		},
			dbx.HashExp{
				"user_id": FID,
			}).Execute()

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func getSaltUser(FID string) (string, error) {
	user := entityUser.User{}
	db := dbcontext.GetDB()
	err := db.Select("id").From("users").Where(dbx.HashExp{"front_id": FID}).One(&user)
	if err != nil {
		return "", err
	}

	profile := entityUser.Profile{}
	err = db.Select("salt").From("profiles").Where(dbx.HashExp{"user_id": user.ID}).One(&profile)
	if err != nil {
		return "", err
	}
	return profile.Salt, nil
}

func countProfileInDB(newUser entityUser.UserCreate) (int, error) {
	q := dbcontext.GetDB().NewQuery("SELECT user_id FROM profiles WHERE login={:login} OR email={:email} OR phone={:phone}")
	q.Bind(dbx.Params{
		"login": newUser.Login,
		"email": newUser.Email,
		"phone": newUser.Phone,
	})
	profile := entityUser.Profile{}
	err := q.One(&profile)

	if err != nil {
		return -0, err
	}

	var users []*entityUser.User
	err = q.All(users)

	if err != nil {
		return -0, err
	}

	return len(users), nil
}

func isUserInDB(FID string) error {
	newUser := entityUser.User{}

	q := dbcontext.GetDB().NewQuery("SELECT id FROM users WHERE front_id={:id}")
	q.Bind(dbx.Params{"id": FID})

	err := q.One(&newUser)
	if err != nil {
		return err
	}

	user := entityUser.User{}
	err = q.Row(&user)
	if err != nil {
		return nil
	}

	profile := entityUser.Profile{}
	q = dbcontext.GetDB().NewQuery("SELECT id FROM profiles WHERE user_id={:id}")
	q.Bind(dbx.Params{"id": user.ID})

	err = q.One(&profile)
	if err != nil {
		return err
	}

	return nil
}

func deleteUserFromStorage(FID string) error {
	err := dbcontext.GetDB().Transactional(func(tx *dbx.Tx) error {
		var err error

		user := entityUser.User{}
		q := dbcontext.GetDB().NewQuery("SELECT id FROM users WHERE front_id={:id}")
		q.Bind(dbx.Params{"id": FID})

		err = q.One(&user)
		if err != nil {
			return err
		}

		// soft deleting from USERS
		_, err = tx.Update("users", dbx.Params{
			"is_deleted": true,
		},
			dbx.HashExp{"id": FID},
		).Execute()

		if err != nil {
			return err
		}

		// soft deleting from PROFILE
		_, err = tx.Update("profiles", dbx.Params{
			"is_deleted": true,
		},
			dbx.HashExp{"user_id": user.ID},
		).Execute()

		if err != nil {
			return err
		}

		// soft deleting from PASSWORDS HISTORY
		_, err = tx.Update("users-passwords-history", dbx.Params{
			"is_deleted": true,
		},
			dbx.HashExp{"user_id": user.ID},
		).Execute()

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func userByFIDFromStorage(FID uuid.UUID) (entityUser.User, error) {
	foundUser := entityUser.User{}

	err := dbcontext.GetDB().Select().Where(dbx.HashExp{"front_id": FID}).One(&foundUser)
	if err != nil {
		return foundUser, err
	}

	return foundUser, nil
}
