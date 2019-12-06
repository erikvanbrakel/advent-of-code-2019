package pkg

func calculateFuelForMass(mass int) int {
    if mass <= 5 {
        return 0
    }

    return (mass / 3) - 2
}

func CalculateTotalFuel(mass int) int {
    var aggregator func(int, int) int

    aggregator = func(mass, runningTotal int) int {
        additionalFuel := calculateFuelForMass(mass)
        if additionalFuel == 0 {
            return runningTotal
        }

        return aggregator(additionalFuel, runningTotal+additionalFuel)
    }

    return aggregator(mass, 0)
}
