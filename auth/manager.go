package auth

type Manager struct{}

func NewManager() *Manager

func (m *Manager) NewUser(u, p string) (*User, error)

func (m *Manager) GetUser(u, p string) (*User, error)
