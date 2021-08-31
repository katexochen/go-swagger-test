package main

// swagger:response firstResponse
type FirstResponse struct {
	// in:body
	Body struct {
		R First
	}
}

// swagger:response secondResponse
type SecondResponse struct {
	// in:body
	Body struct {
		R Second
	}
}

// swagger:response thirdResponse
type ThirdResponse struct {
	// in:body
	Body struct {
		R Third
	}
}

// swagger:response fourthResponse
type FourthResponse struct {
	// in:body
	Body struct {
		R Fourth
	}
}

// First is described in this line.
// swagger:model
type First struct {
	foo int
}

// Second is described in this line
// swagger:model
type Second struct {
	bar int
}

// Third is described in this line.
type Third struct {
	foo int
}

// Fourth is described in this line
type Fourth struct {
	bar int
}
