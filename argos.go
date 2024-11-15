package argos

import (
	"argos/common"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ValidateMode(mode *string) error {
	if mode == nil || *mode == "" {
		return nil
	}
	if *mode == "car" || *mode == "truck" {
		return nil
	}
	return errors.New("invalid mode")
}

func ValidateHazmatType(hazmatType *string, isFlexible bool) error {
	if hazmatType == nil || *hazmatType == "" {
		return nil
	}
	if !isFlexible {
		return errors.New("only flex can support hazmat type")
	}
	items := strings.Split(*hazmatType, "|")
	for _, v := range items {
		v = strings.ToLower(strings.TrimSpace(v))
		switch v {
		case "general", "harmful_to_water", "explosive", "circumstantial":
			continue
		default:
			return errors.New("invalid hazmat type: " + v)
		}
	}
	return nil
}

func ValidateTruckWeight(truckWeight *uint, isFlexible bool) error {
	if truckWeight == nil {
		return nil
	}
	if !isFlexible {
		return errors.New("only flex can support truck weight")
	}
	v := float64(*truckWeight)
	v = v / 1000.0
	if v > 100.0 {
		return errors.New("invalid truck_weight, should in range [0, 100] tons")
	}
	return nil
}

func ValidateTruckSize(truckSize *string, isFlexible bool) error {
	if truckSize == nil || *truckSize == "" {
		return nil
	}
	if !isFlexible {
		return errors.New("only flex can support truck size")
	}
	items := strings.Split(*truckSize, ",")
	if len(items) != 3 {
		return errors.New("truck size should be 3 items")
	}
	height, err := strconv.ParseFloat(strings.TrimSpace(items[0]), 64)
	if err != nil {
		return err
	}
	height = height / 100.0
	width, err := strconv.ParseFloat(strings.TrimSpace(items[1]), 64)
	if err != nil {
		return err
	}
	width = width / 100.0
	length, err := strconv.ParseFloat(strings.TrimSpace(items[2]), 64)
	if err != nil {
		return err
	}
	length = length / 100.0
	if length < 0 || length > 50 ||
		width < 0 || width > 50 ||
		height < 0 || height > 10 {
		return errors.New("invalid truck_size, should be in range [0, 50] meters for length and width, [0, 10] meters for height")
	}
	return nil
}

func ValidateTruckAxleLoad(truckAxleLoad *float64, isFlexible bool) error {
	if truckAxleLoad == nil {
		return nil
	}
	if !isFlexible {
		return errors.New("only flex can support truck axle load")
	}
	if *truckAxleLoad < 0 {
		return errors.New("invalid truck_axle_load, should be greater than 0")
	}
	return nil
}

func ValidateAvoid(avoid *string, isFlexible bool) error {
	if avoid == nil || *avoid == "" || *avoid == "none" {
		return nil
	}
	if !isFlexible {
		items := strings.Split(*avoid, "|")
		for _, v := range items {
			v = strings.ToLower(strings.TrimSpace(v))
			switch v {
			case "highway", "highways", "motorway", "motorways", "toll", "tolls", "ferry", "ferries":
				continue
			default:
				return errors.New("invalid avoid type: " + v)
			}
		}
	} else {
		items := strings.Split(*avoid, "|")
		for _, v := range items {
			v = strings.ToLower(strings.TrimSpace(v))
			switch v {
			case "left_turn", "right_turn", "single_lane", "highway", "highways", "motorway", "motorways",
				"toll", "tolls", "ferry", "ferries", "uturn", "uturns", "sharp_turn", "sharp_turns",
				"living_street", "living_streets", "service_road", "service_roads":
				continue
			default:
				if strings.HasPrefix(v, "bbox:") {
					_, err := common.ParseAvoidBBox(v)
					if err != nil {
						return err
					}
				} else if strings.HasPrefix(v, "max_speed:") {
					_, max_speed_str, _ := strings.Cut(v, "max_speed:")
					max_speed, err := strconv.ParseFloat(max_speed_str, 64)
					if err != nil {
						return errors.New("invalid max_speed")
					}
					if max_speed < 0 {
						return errors.New("invalid max_speed, should be greater than 0")
					}
				} else {
					return errors.New("unsupport avoid object")
				}
			}
		}
	}
	return nil
}

func ValidateApproaches(approaches *string, pointsNum int) error {
	if approaches == nil || *approaches == "" {
		return nil
	}
	da, err := common.ParseApproaches(approaches)
	if err != nil {
		return err
	}
	if len(da) != pointsNum {
		return fmt.Errorf("the number of approaches should be %v", pointsNum)
	}
	return nil
}