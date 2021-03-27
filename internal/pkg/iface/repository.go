package iface

type IRepository interface {
	CardRepo
	UserRepo
}

type CardRepo interface {
}

type UserRepo interface {
}
