package entity

type (
	Collider interface {
		IsCollideWith(other Collider) bool
	}
)
