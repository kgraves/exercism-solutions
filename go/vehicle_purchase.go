package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	ret := option1
    if option2 < option1 {
        ret = option2
    }
    return ret + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    factor := 0.8
    if age >= 3 && age < 10 {
        factor = 0.7
    } else if age >= 10 {
        factor = 0.5
    }

    return factor * originalPrice
}
