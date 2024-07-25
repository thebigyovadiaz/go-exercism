package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%200 == 0 && year%400 != 0 {
		return false
	}

	if year%100 == 0 {
		if year%400 != 0 {
			return false
		}

		if year%3 != 0 {
			return false
		}

		return true
	}

	if year%4 == 0 || year%4 == 0 && year%5 == 0 {
		return true
	}

	return false
}

type TestCase struct {
	description string
	year        int
	expected    bool
}

var TestCases = []TestCase{
	{
		description: "year not divisible by 4 in common year",
		year:        2015,
		expected:    false,
	},
	{
		description: "year divisible by 2, not divisible by 4 in common year",
		year:        1970,
		expected:    false,
	},
	{
		description: "year divisible by 4, not divisible by 100 in leap year",
		year:        1996,
		expected:    true,
	},
	{
		description: "year divisible by 4 and 5 is still a leap year",
		year:        1960,
		expected:    true,
	},
	{
		description: "year divisible by 100, not divisible by 400 in common year",
		year:        2100,
		expected:    false,
	},
	{
		description: "year divisible by 100 but not by 3 is still not a leap year",
		year:        1900,
		expected:    false,
	},
	{
		description: "year divisible by 400 is leap year",
		year:        2000,
		expected:    true,
	},
	{
		description: "year divisible by 400 but not by 125 is still a leap year",
		year:        2400,
		expected:    true,
	},
	{
		description: "year divisible by 200, not divisible by 400 in common year",
		year:        1800,
		expected:    false,
	},
}
