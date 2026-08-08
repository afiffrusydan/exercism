package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate)*successRate/float64(100.0))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate)/float64(60.0)*successRate/float64(100.0))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	modulo := uint(carsCount)%uint(10)
    puluhan := (uint(carsCount)-modulo)/uint(10)
    return  puluhan*uint(95000)+modulo*uint(10000)
}
