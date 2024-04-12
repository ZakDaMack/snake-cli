package models

type Snake struct {
	body      []Position
	bounds    Position
	length    int
	direction int
}

const (
	START_LENGTH    = 2
	DIRECTION_UP    = 0
	DIRECTION_RIGHT = 1
	DIRECTION_DOWN  = 2
	DIRECTION_LEFT  = 3
)

func NewSnake(upperBounds Position) *Snake {
	return &Snake{
		body: []Position{
			{upperBounds.X / 2, upperBounds.Y / 2},
		},
		length:    START_LENGTH,
		direction: DIRECTION_RIGHT,
		bounds:    upperBounds,
	}
}

func (s *Snake) Move() []Position {
	s.addHead()
	s.trim()
	return s.body
}

func (s *Snake) IncreaseLength() {
	s.length++
}

func (s *Snake) ChangeDirection(direction int) {
	s.direction = direction
}

func (s *Snake) addHead() {
	p := s.body[len(s.body)-1]
	switch s.direction {
	case DIRECTION_UP:
		p.Y++
	case DIRECTION_RIGHT:
		p.X++
	case DIRECTION_DOWN:
		p.Y--
	case DIRECTION_LEFT:
		p.X--
	}

	// if the snake goes out of bounds, wrap around
	p.X = p.X % s.bounds.X
	p.Y = p.Y % s.bounds.Y
	s.body = append(s.body, p)
}

func (s *Snake) trim() {
	if len(s.body) > s.length {
		s.body = s.body[1:]
	}
}
