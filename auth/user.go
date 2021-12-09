package auth

type User struct {
	ID       int
	Username string
	Password string
	Enabled  bool
	Locked   bool
}

func (u *User) Login(p string) error

func (u *User) Enable() error

func (u *User) Disable() error

func (u *User) Lock() error

func (u *User) Unlock() error
