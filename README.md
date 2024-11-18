# Argos
## Argos Validation Functions
This module contains several functions to validate navigation parameters, ensuring the parameters are valid when provided to Route Engine Service.

## Function Descriptions For Navigation API

### 1. `ValidateMode(mode *string) error`

Validates the vehicle mode. Allowed modes are `"car"` and `"truck"`.

#### Parameters:
- `mode`: A pointer to a string representing the vehicle mode.

#### Return:
- `nil`: If the mode is valid or empty.
- `Error`: If the mode is invalid.

### 2. `ValidateHazmatType(hazmatType *string, isFlexible bool) error`

Validates the hazmat type. If `isFlexible` is `false` (OSRM engine), hazmat types are not supported.

#### Parameters:
- `hazmatType`: A pointer to a string representing the hazmat type.
- `isFlexible`: A boolean indicating whether the engine is flexible (Valhalla engine).

#### Return:
- `nil`: If the hazmat type is valid.
- `Error`: If the hazmat type is invalid or not supported by the OSRM engine.

### 3. `ValidateTruckWeight(truckWeight *uint, isFlexible bool) error`

Validates the truck weight. If `isFlexible` is `false` (OSRM engine), truck weight is not supported.

#### Parameters:
- `truckWeight`: A pointer to an unsigned integer representing the truck weight (in kilograms).
- `isFlexible`: A boolean indicating whether the engine is flexible (Valhalla engine).

#### Return:
- `nil`: If the truck weight is valid.
- `Error`: If the truck weight exceeds the limit or if the OSRM engine doesn't support it.

### 4. `ValidateTruckSize(truckSize *string, isFlexible bool) error`

Validates the truck size. If `isFlexible` is `false` (OSRM engine), truck size is not supported.

#### Parameters:
- `truckSize`: A pointer to a string representing the truck size (format: height,width,length in centimeters).
- `isFlexible`: A boolean indicating whether the engine is flexible (Valhalla engine).

#### Return:
- `nil`: If the truck size is valid.
- `Error`: If the truck size format is incorrect, exceeds the range, or is not supported by the OSRM engine.

### 5. `ValidateTruckAxleLoad(truckAxleLoad *float64, isFlexible bool) error`

Validates the truck axle load. If `isFlexible` is `false` (OSRM engine), truck axle load is not supported.

#### Parameters:
- `truckAxleLoad`: A pointer to a float64 representing the truck axle load (in kilograms).
- `isFlexible`: A boolean indicating whether the engine is flexible (Valhalla engine).

#### Return:
- `nil`: If the truck axle load is valid.
- `Error`: If the axle load is invalid or if the OSRM engine doesn't support it.

### 6. `ValidateAvoid(avoid *string, isFlexible bool) error`

Validates the `avoid` parameter, which specifies the road types to avoid. If `isFlexible` is `false` (OSRM engine), only basic avoid options are supported.

#### Parameters:
- `avoid`: A pointer to a string representing the road types to avoid.
- `isFlexible`: A boolean indicating whether the engine is flexible (Valhalla engine).

#### Return:
- `nil`: If the avoid types are valid.
- `Error`: If the avoid types are invalid or not supported by the engine.

### 7. `ValidateApproaches(approaches *string, pointsNum int) error`

Validates the `approaches` parameter by checking if the number of points matches the expected number.

#### Parameters:
- `approaches`: A pointer to a string representing a list of approach points.
- `pointsNum`: The expected number of approach points.

#### Return:
- `nil`: If the number of approach points matches the expected value.
- `Error`: If the number of approach points does not match, or if the `approaches` string contains invalid characters.


## Error Handling

These functions return an `error` type value. If an input does not meet the requirements, the function returns an error containing the appropriate error message. If the input is valid, the function returns `nil`.
