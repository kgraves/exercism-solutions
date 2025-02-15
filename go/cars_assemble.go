package cars

// CarCost is the cost to produce an individual car.
const CarCost = 10_000
// GroupCost is the cost to product a group of 10 cats together.
const GroupCost = 95_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    groups := uint(carsCount / 10)
    singles := uint(carsCount % 10)

    return (groups * GroupCost) + (singles * CarCost)
}
