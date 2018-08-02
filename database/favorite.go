package database

type Favorite struct {
	DeviceID string `storm:"index"`
}
