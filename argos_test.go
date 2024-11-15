package argos

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMode(t *testing.T) {
	tests := []struct {
		mode     *string
		expected error
	}{
		{nil, nil},
		{stringPtr(""), nil},
		{stringPtr("car"), nil},
		{stringPtr("truck"), nil},
		{stringPtr("bike"), errors.New("invalid mode")},
	}

	for _, test := range tests {
		var name string
		if test.mode == nil {
			name = "nil"
		} else {
			name = *test.mode
		}
		t.Run(name, func(t *testing.T) {
			err := ValidateMode(test.mode)
			if test.expected != nil {
				assert.Equal(t, test.expected.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateHazmatType(t *testing.T) {
	tests := []struct {
		hazmatType *string
		isFlexible bool
		expected   error
	}{
		{nil, true, nil},
		{stringPtr(""), true, nil},
		{stringPtr("general"), true, nil},
		{stringPtr("explosive|circumstantial"), true, nil},
		{stringPtr("general|unknown"), true, errors.New("invalid hazmat type: unknown")},
		{stringPtr("explosive"), false, errors.New("only flex can support hazmat type")},
	}

	for _, test := range tests {
		var name string
		if test.hazmatType == nil {
			name = "nil"
		} else {
			name = *test.hazmatType
		}
		t.Run(name, func(t *testing.T) {
			err := ValidateHazmatType(test.hazmatType, test.isFlexible)
			if test.expected != nil {
				assert.Equal(t, test.expected.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateTruckWeight(t *testing.T) {
	tests := []struct {
		truckWeight *uint
		isFlexible  bool
		expected    error
	}{
		{nil, true, nil},
		{uintPtr(5000), true, nil}, // 5 tons
		{uintPtr(150000), true, errors.New("invalid truck_weight, should in range [0, 100] tons")},
		{uintPtr(5000), false, errors.New("only flex can support truck weight")},
	}

	for _, test := range tests {
		var name string
		if test.truckWeight == nil {
			name = "nil"
		} else {
			name = strconv.Itoa(int(*test.truckWeight))
		}
		t.Run(name, func(t *testing.T) {
			err := ValidateTruckWeight(test.truckWeight, test.isFlexible)
			if test.expected != nil {
				assert.Equal(t, test.expected.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateTruckSize(t *testing.T) {
	tests := []struct {
		truckSize  *string
		isFlexible bool
		expected   error
	}{
		{nil, true, nil},
		{stringPtr("400, 300, 800"), true, nil},
		{stringPtr("120, 300, 5100"), true, errors.New("invalid truck_size, should be in range [0, 50] meters for length and width, [0, 10] meters for height")},
		{stringPtr("400, 300, 800"), false, errors.New("only flex can support truck size")},
	}

	for _, test := range tests {
		var name string
		if test.truckSize == nil {
			name = "nil"
		} else {
			name = *test.truckSize
		}
		t.Run(name, func(t *testing.T) {
			err := ValidateTruckSize(test.truckSize, test.isFlexible)
			if test.expected != nil {
				if err != nil {
					assert.Equal(t, test.expected.Error(), err.Error())
				} else {
					t.Errorf("expected error but got nil")
				}
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateTruckAxleLoad(t *testing.T) {
	tests := []struct {
		truckAxleLoad *float64
		isFlexible    bool
		expected      error
	}{
		{nil, true, nil},
		{float64Ptr(1000), true, nil},
		{float64Ptr(-500), true, errors.New("invalid truck_axle_load, should be greater than 0")},
		{float64Ptr(1000), false, errors.New("only flex can support truck axle load")},
	}

	for _, test := range tests {
		var name string
		if test.truckAxleLoad == nil {
			name = "nil"
		} else {
			name = strconv.FormatFloat(*test.truckAxleLoad, 'f', -1, 64)
		}
		t.Run(name, func(t *testing.T) {
			err := ValidateTruckAxleLoad(test.truckAxleLoad, test.isFlexible)
			if test.expected != nil {
				assert.Equal(t, test.expected.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateAvoid(t *testing.T) {
	tests := []struct {
		avoid      *string
		isFlexible bool
		expected   error
	}{
		{nil, true, nil},
		{stringPtr("none"), true, nil},
		{stringPtr("highway|ferry"), true, nil},
		{stringPtr("unknown"), true, errors.New("error")},
		{stringPtr("highway|ferry"), false, nil},
		{stringPtr("unknown|ferry"), false, errors.New("error")},
	}

	for _, test := range tests {
		var name string
		if test.avoid == nil {
			name = "nil"
		} else {
			name = *test.avoid
		}
		t.Run(name, func(t *testing.T) {
			err := ValidateAvoid(test.avoid, test.isFlexible)
			if test.expected != nil {
				// 不需要判断err的内容 仅仅判断是否有错误
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateApproaches(t *testing.T) {
	tests := []struct {
		approaches *string
		pointsNum  int
		expected   error
	}{
		{nil, 2, nil},
		{stringPtr(";"), 2, nil},
		{stringPtr("curb;curb"), 3, fmt.Errorf("the number of approaches should be %v", 3)},
		{stringPtr("curb;unrestricted"), 2, nil},
		{stringPtr("curb;unrestricted;;"), 4, nil},
	}

	for _, test := range tests {
		var name string
		if test.approaches == nil {
			name = "nil"
		} else {
			name = *test.approaches
		}
		t.Run(name, func(t *testing.T) {
			err := ValidateApproaches(test.approaches, test.pointsNum)
			if test.expected != nil {
				assert.Equal(t, test.expected.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func uintPtr(u uint) *uint {
	return &u
}

func float64Ptr(f float64) *float64 {
	return &f
}
