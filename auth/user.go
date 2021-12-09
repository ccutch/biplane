package auth

type User struct {
	manager *Manager

	ID       int
	Username string
	Password string
	Enabled  bool
	Locked   bool
}

func (u *User) Enable() error {
	sql := "update users set enabled = true where id = $1"

	_, err := u.manager.Config.Client().Exec(sql, u.ID)
	if err != nil {
		u.Enabled = true
	}

	return err
}

func (u *User) Disable() error {
	sql := "update users set enabled = false where id = $1"

	_, err := u.manager.Config.Client().Exec(sql, u.ID)
	if err != nil {
		u.Enabled = false
	}

	return err
}

func (u *User) Lock() error {
	sql := "update users set locked = true where id = $1"

	_, err := u.manager.Config.Client().Exec(sql, u.ID)
	if err != nil {
		u.Locked = true
	}

	return err
}

func (u *User) Unlock() error {
	sql := "update users set locked = false where id = $1"

	_, err := u.manager.Config.Client().Exec(sql, u.ID)
	if err != nil {
		u.Locked = false
	}

	return err
}
