package two

/* pt1

o p
A X Rock 1
B Y Paper 2
C Z Scissors 3

L 0
D 3
W 6

*/

/* pt2

o p
A X Rock 1
B Y Paper 2
C Z Scissors 3

L 0
D 3
W 6

*/

type Shape struct {
    Rock     string
    Paper    string
    Scissors string
}

type ShapeRules struct {
    Rock     int
    Paper    int
    Scissors int
}


type Rules struct {
    Loss int
    Draw int
    Win  int
}

var shapeRules = ShapeRules{
	Rock:     1,
    Paper:    2,
    Scissors: 3,
}

var ruleSet = Rules{
	Loss: 0,
	Draw: 3,
	Win:  6,
}