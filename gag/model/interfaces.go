package model

import "context"

// handler layer
type DeviceService interface {
	Get(ctx context.Context, uuid string) (*Device, error)
	Register(ctx context.Context, d *Device) error
	Delete(ctx context.Context, uuid string) error
}

type UserService interface {
	Get(id string) (*User, error)
	Login(ctx context.Context, u *User) error
}

// repository layer
type DeviceRepository interface {
	FindByID(uuid string) (*Device, error)
	Create(ctx context.Context, d *Device) error
	Delete(ctx context.Context, uuid string) error
}

type UserRepository interface {
	FindByID(id string) (*User, error)
	Create(ctx context.Context, u *User) error
}
