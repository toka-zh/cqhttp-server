package module

type IModule interface {
	Enable() error
	Disable() error
	Status() string
}
