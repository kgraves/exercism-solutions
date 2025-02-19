package interest

const NegativeBalanceInterest = 3.213
const SubOneThousandInterest = 0.5
const SubFiveThousandInterest = 1.621
const OverFiveThousandInterest = 2.475

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    switch {
    case balance < 0.0:
        return NegativeBalanceInterest
    case balance < 1000.0:
        return SubOneThousandInterest
    case balance < 5000.0:
        return SubFiveThousandInterest
    case balance >= 5000.0:
        return OverFiveThousandInterest
    default:
        panic("ruh roh, unexpected balance value")
    }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance) / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    years := 0

	for years = 0; balance < targetBalance; years++ {
        balance += Interest(balance)
    }

    return years
}
