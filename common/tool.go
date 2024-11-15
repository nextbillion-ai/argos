package common

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

var ErrAvoidBBOXFormat = errors.New("avoid bbox format should be `avoid=bbox:min_lat,min_lon,max_lat,max_lon`")

func (c *Coordinate) Valid() error {
	if c == nil {
		return errors.New("nil pointer of coordinate")
	}
	if c.Lat > 90.0 || c.Lat < -90.0 {
		return fmt.Errorf("invalid coordinate, latitude: %v", c.Lat)
	}

	if c.Lon > 180.0 || c.Lon < -180.0 {
		return fmt.Errorf("invalid coordinate, longitude: %v", c.Lon)
	}
	return nil
}

func parseFloatNoNaN(v string) (float64, error) {
	lat, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
	if err != nil {
		return 0.0, err
	}

	if math.IsNaN(lat) {
		return 0.0, errors.New("NaN is not allowed as input")
	}

	return lat, nil
}

func ParseAvoidBBox(avoid string) ([]*Coordinate, error) {
	_, v, _ := strings.Cut(avoid, "bbox:")
	bbox := strings.Split(v, ",")
	if len(bbox) != 4 {
		return nil, ErrAvoidBBOXFormat
	}

	min_lat, err := parseFloatNoNaN(bbox[0])
	if err != nil {
		return nil, ErrAvoidBBOXFormat
	}
	min_lon, err := parseFloatNoNaN(bbox[1])
	if err != nil {
		return nil, ErrAvoidBBOXFormat
	}
	max_lat, err := parseFloatNoNaN(bbox[2])
	if err != nil {
		return nil, ErrAvoidBBOXFormat
	}
	max_lon, err := parseFloatNoNaN(bbox[3])
	if err != nil {
		return nil, ErrAvoidBBOXFormat
	}
	if min_lat > max_lat || min_lon > max_lon {
		return nil, ErrAvoidBBOXFormat
	}
	min_coord := &Coordinate{
		Lat: min_lat,
		Lon: min_lon,
	}
	if err := min_coord.Valid(); err != nil {
		return nil, ErrAvoidBBOXFormat
	}
	max_coord := &Coordinate{
		Lat: max_lat,
		Lon: max_lon,
	}
	if err := max_coord.Valid(); err != nil {
		return nil, ErrAvoidBBOXFormat
	}
	return []*Coordinate{
		min_coord,
		{
			Lat: min_lat,
			Lon: max_lon,
		},
		max_coord,
		{
			Lat: max_lat,
			Lon: min_lon,
		},
	}, nil
}

type PreferredSide string

const (
	Same     PreferredSide = "same"
	Opposite PreferredSide = "opposite"
	Either   PreferredSide = "either"
)

func ParseApproaches(approaches *string) ([]PreferredSide, error) {
	if approaches == nil || *approaches == "" {
		return nil, nil
	}
	items := strings.Split(*approaches, ";")
	res := make([]PreferredSide, 0, len(items))
	for _, v := range items {
		v = strings.ToLower(strings.TrimSpace(v))
		switch v {
		case "curb":
			res = append(res, Same)
		case "", "unrestricted":
			res = append(res, Either)
		default:
			return nil, errors.New("invalid approach")
		}
	}
	return res, nil
}
