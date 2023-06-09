package questions

import (
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

const (
	BiologicalScienceDegree = "Biological Science"
	ComputerScienceDegree   = "Computer Science"
	EngineeringDegree       = "Engineering"
	MathematicsDegree       = "Mathematics"
	SocialScienceDegree     = "Social Science"
	OtherDegrees            = "Other"
)

var Departments = []string{
	BiologicalScienceDegree,
	ComputerScienceDegree,
	EngineeringDegree,
	MathematicsDegree,
	SocialScienceDegree,
	OtherDegrees,
}

const (
	AstronautIncurableEyeSightProblem = "Yes, I have incurable eyesight problem"
	AstronautCurableEyeSightProblem   = "Yes, I have eyesight problem but it is curable"
	AstronautNoEyeSightProblem        = "No, I have no eyesight problem"
)

var EyeSightOptions = []string{
	AstronautIncurableEyeSightProblem,
	AstronautCurableEyeSightProblem,
	AstronautNoEyeSightProblem,
}

type AstronautAnswer struct {
	Tall           int
	EyeSight       Select
	BachelorDegree Select
	MasterDegree   Select
}

var AstronautValidDegree = map[string]bool{
	BiologicalScienceDegree: true,
	ComputerScienceDegree:   true,
	EngineeringDegree:       true,
	MathematicsDegree:       true,
}

const (
	AstronautMinRestrictionHeight = 149
	AstronautMaxRestrictionHeight = 193
)

var AstronautQuestions = []*survey.Question{
	{
		Name:      "Tall",
		Prompt:    &survey.Input{Message: "How tall are you? (cm)"},
		Validate:  survey.Required,
		Transform: transformInt,
	},
	{
		Name: "BachelorDegree",
		Prompt: &survey.Select{
			Message: "What is you degree?",
			Options: Departments,
		},
		Validate: survey.Required,
	},
	{
		Name: "MasterDegree",
		Prompt: &survey.Select{
			Message: "What is your master degree/s?",
			Options: Departments,
		},
		Validate: survey.Required,
	},
}

var transformInt = func(ans interface{}) (newAns interface{}) {
	strValue, ok := ans.(string)
	if !ok {
		return 0
	}
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0
	}

	return intValue

}
