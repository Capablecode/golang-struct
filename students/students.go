package students

// Define the struct for student
type Student struct {
	Name  string
	Age   int
	Score []float64
}

// Adding a method for the student struct, that calculate the average score
func (s Student) AverageScore() float64 {
	if len(s.Score) == 0 {
		return 0
	}
	var total float64

	for _, score := range s.Score {
		total += score
	}
	return total / float64(len(s.Score))
}

// A method to determine if student passed or failed
func (s Student) HasPassed() bool {
	currentAvaerage := s.AverageScore()
	return currentAvaerage >= 50
}