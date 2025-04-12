package models

type Migration struct{}

func (m *Migration) RegisterModels() []interface{} {
    return []interface{}{
        &User{},
    }
}