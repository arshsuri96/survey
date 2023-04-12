package strategy

import (
	"arshsuri96/JOB/questions"
	"encoding/json"

	"github.com/AlecAivazis/survey/v2"
)

type astronaut struct{}

const (
	AstronautInappropriateMsgText = "Sorry, There are really strict rules to be selected. Better to prefer another job."
	AstronautAppropriateMsgText   = "You have met the pre-conditions. If you selected, you need to work hard :)"
)

func (ast *astronaut) Check(answerBytes []byte) (CheckResult, error) {
	var answer questions.AstronautAnswer
	err := json.Unmarshal(answerBytes, &answer)
	if err != nil {
		return CheckResult{}, err
	}

	if ok := questions.AstronautValidDegree[answer.BachelorDegree.Value.(string)]; !ok {
		return CheckResult{text: AstronautInappropriateMsgText, Status: false}, nil
	}
	if ok := questions.AstronautValidDegree[answer.MasterDegree.Value.(string)]; !ok {
		return CheckResult{text: AstronautInappropriateMsgText, Status: false}, nil
	}

	eyeSightVal := answer.EyeSight.Value

	if (eyeSightVal == questions.AstronautNoEyeSightProblem || eyeSightVal == questions.AstronautCurableEyeSightProblem) && (answer.Tall > questions.AstronautMinRestrictionHeight && answer.Tall < questions.AstronautMaxRestrictionHeight) {
		return CheckResult{text: AstronautAppropriateMsgText, Status: true}, nil
	}

	return CheckResult{text: AstronautAppropriateMsgText, Status: false}, nil
}

func (ast *astronaut) GetQuestions() []*survey.Question {
	return questions.AstronautQuestions
}
