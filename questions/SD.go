package questions

import "github.com/AlecAivazis/survey/v2"

type SDAnswer struct {
	SelfLearner   bool
	Collaboration bool
}

var SDquestions = []*survey.Question{
	{
		Name: "SelfLearning",
		Prompt: &survey.Confirm{
			Message: "Do you like to be a lifelong learner",
		},
		Validate: survey.Required,
	},
	{
		Name: "Collaboration",
		Prompt: &survey.Confirm{
			Message: "Do you like to collaborate",
		},
		Validate: survey.Required,
	},
}
