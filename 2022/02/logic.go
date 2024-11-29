package two

func rock(p string)int{
	switch p {
		case pShape.Rock:
			return ruleSet.Draw + shapeRules.Rock
		case pShape.Paper:
			return ruleSet.Win + shapeRules.Paper
		case pShape.Scissors:
			return ruleSet.Loss + shapeRules.Scissors
		default:
			return 0;
	}
	
}

func paper(p string)int{
	switch p {
		case pShape.Rock:
			return ruleSet.Loss + shapeRules.Rock
		case pShape.Paper:
			return ruleSet.Draw + shapeRules.Paper
		case pShape.Scissors:
			return ruleSet.Win + shapeRules.Scissors
		default:
			return 0;
	}
}

func scissors(p string)int{
	switch p {
		case pShape.Rock:
			return ruleSet.Win + shapeRules.Rock
		case pShape.Paper:
			return ruleSet.Loss + shapeRules.Paper
		case pShape.Scissors:
			return ruleSet.Draw + shapeRules.Scissors
		default:
			return 0;
	}
}