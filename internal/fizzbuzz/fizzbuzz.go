package fizzbuzz

import (
	"errors"
	"strconv"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateSequence(string1 string, string2 string, int1 int, int2 int, limit int) ([]string, error) {
	if limit <= 0 {
		return nil, errors.New("limit must be positive")
	}
	if int1 <= 0 || int2 <= 0 {
		return nil, errors.New("int1 and int2 must be positive")
	}
	if string1 == "" || string2 == "" {
		return nil, errors.New("string1 and string2 must be non-empty")
	}
	if len(string1) > 1000 || len(string2) > 1000 {
		return nil, errors.New("string1 and string2 must be less than 1000 characters")
	}

	seq := make([]string, limit)
	fizzBuzz := string1 + string2

	for i := 1; i <= limit; i++ {
		fizz := i%int1 == 0
		buzz := i%int2 == 0
		switch {
		case fizz && buzz:
			seq[i-1] = fizzBuzz
		case fizz:
			seq[i-1] = string1
		case buzz:
			seq[i-1] = string2
		default:
			seq[i-1] = strconv.Itoa(i)
		}
	}

	return seq, nil
}
