package iface

import "gitlab.com/howmay/gopher/db"

type IRepository interface {
	db.IBaseRepository
}
