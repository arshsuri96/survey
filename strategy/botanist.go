package strategy

import (
	"arshsuri96/JOB/questions"
	"encoding/json"

	"github.com/AlecAivazis/survey/v2"
)

type botanist struct{}

const (
	BotanistInappropriateMessage = "Becoming botanist really needs strong interest and passion. You need to whip these." +
		"Don't lose your hope."
	BotanistAppropriateMessageForBot = "You are ideal candidate, however you need to work hard to be Botanist."
)

func (b *botanist) Check(answerBytes []byte) (CheckResult, error) {
	var answer questions.BotanistAnswer
	err := json.Unmarshal(answerBytes, &answer)
	if err != nil {
		return CheckResult{}, nil
	}
	if answer.LabExercise && answer.LocalFlora && answer.NewPlantSpecies && answer.Passion {
		return CheckResult{text: BotanistAppropriateMessageForBot, Status: true}, nil
	}

	return CheckResult{text: BotanistInappropriateMessage, Status: false}, nil
}

func (b *botanist) GetQuestions() []*survey.Question {
	return questions.BotanistQuestions
}
