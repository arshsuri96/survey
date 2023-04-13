package strategy

import (
	"arshsuri96/JOB/questions"
	"encoding/json"

	"github.com/AlecAivazis/survey/v2"
)

type SD struct{}

const (
	SoftwareDeveloperInappropriateMessage = "You need to change your mindset. Don't lose your hope. Try and try again."
	SoftwareDeveloperAppropriateMessage   = "You are ideal candidate, however you need to work hard to be Software Developer"
)

func (s *SD) Check(answerByte []byte) (CheckResult, error) {
	var answer questions.SDAnswer
	err := json.Unmarshal(answerByte, &answer)
	if err != nil {
		return CheckResult{}, nil
	}

	if answer.SelfLearner && answer.Collaboration {
		return CheckResult{text: SoftwareDeveloperAppropriateMessage, Status: true}, nil
	}

	return CheckResult{text: SoftwareDeveloperInappropriateMessage, Status: false}, nil

}

func (s *SD) GetQuestions() []*survey.Question {
	return questions.SDquestions
}
