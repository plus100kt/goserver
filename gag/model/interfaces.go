package model

import "context"

type UserService interface {
	Get(ctx context.Context, id string) (*User, error)
	DeviceRegister(ctx context.Context, uuid string) (*Device, error)
	Login(ctx context.Context, key string, u *User) (*User, error)
}

// repository layer
type DeviceRepository interface {
	FindByID(ctx context.Context, uuid string) (*Device, error)
	Create(ctx context.Context, d *Device) error
	Delete(ctx context.Context, uuid string) error
}

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, u *User) error
}

type EclassRepository interface {
	Login(ctx context.Context, key string, u *User) (*User, error)
}
