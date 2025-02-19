package lasagna

const DefaultPrepTime int = 2

func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = DefaultPrepTime
    }

    return len(layers) * avgPrepTime
}

func Quantities(layers []string) (int, float64) {
    noodlesNeeded := 0
    sauceNeeded := 0.0

    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodlesNeeded += 50
        } else if layers[i] == "sauce" {
            sauceNeeded += 0.2
        }
    }

    return noodlesNeeded, sauceNeeded
}

func AddSecretIngredient(theirs []string, ours []string) {
    ours[len(ours) - 1] = theirs[len(theirs) - 1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaledQuantities := make([]float64, len(quantities))

	for i, v := range quantities {
        scaledQuantities[i] = (v / 2.0) * float64(portions)
    }

    return scaledQuantities
}
