package ginper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	ErrInvalidLatitude     = errors.New("invalid latitude")
	ErrInvalidLongitude    = errors.New("invalid longitude")
	ErrLatitudeOutOfRange  = errors.New("latitude out of range")
	ErrLongitudeOutOfRange = errors.New("longitude out of range")
)

type LatLong struct {
	Lat  float64
	Long float64
}

func ParseLatLong(c *gin.Context) (LatLong, error) {
	lat := c.Query("lat")
	long := c.Query("long")

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return LatLong{}, ErrInvalidLatitude
	}

	longitude, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return LatLong{}, ErrInvalidLongitude
	}

	if latitude > 90 || latitude < -90 {
		return LatLong{}, ErrLatitudeOutOfRange
	}

	if longitude > 180 || longitude < -180 {
		return LatLong{}, ErrLongitudeOutOfRange
	}
	return LatLong{Lat: latitude, Long: longitude}, nil
}
