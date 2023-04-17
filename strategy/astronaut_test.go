package strategy

import (
	"encoding/json"
	"testing"

	"arshsuri96/JOB/questions"

	"github.com/stretchr/testify/assert"
)

func TestAstronaut_Check(t *testing.T) {
	t.Run("Appropriate_Answers", func(t *testing.T) {
		// Given
		astronaut := astronaut{}
		answer := questions.AstronautAnswer{
			Tall: 165,
			EyeSight: questions.Select{
				Index: 2,
				Value: questions.AstronautNoEyeSightProblem,
			},
			BachelorDegree: questions.Select{
				Index: 0,
				Value: questions.BiologicalScienceDegree,
			},
			MasterDegree: questions.Select{
				Index: 0,
				Value: questions.BiologicalScienceDegree,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "You have met the pre-conditions. If you selected, you need to work hard :)"
		// Then
		assert.Equal(t, expectedText, check.text)
		assert.True(t, check.Status)
	})
	t.Run("When_Tall_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		astronaut := astronaut{}
		answer := questions.AstronautAnswer{
			Tall: 120,
			EyeSight: questions.Select{
				Index: 0,
				Value: questions.AstronautNoEyeSightProblem,
			},
			BachelorDegree: questions.Select{
				Index: 0,
				Value: questions.BiologicalScienceDegree,
			},
			MasterDegree: questions.Select{
				Index: 0,
				Value: questions.BiologicalScienceDegree,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "Sorry, There are really strict rules to be selected. Better to prefer another job."

		//Then
		assert.Equal(t, expectedText, check.text)
		assert.False(t, check.Status)
	})
	t.Run("When_Bachelor_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		astronaut := astronaut{}
		answer := questions.AstronautAnswer{
			Tall: 120,
			EyeSight: questions.Select{
				Index: 0,
				Value: questions.AstronautCurableEyeSightProblem,
			},
			BachelorDegree: questions.Select{
				Index: 0,
				Value: questions.SocialScienceDegree,
			},
			MasterDegree: questions.Select{
				Index: 0,
				Value: questions.BiologicalScienceDegree,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "Sorry, There are really strict rules to be selected. Better to prefer another job."

		//Then
		assert.Equal(t, expectedText, check.text)
		assert.False(t, check.Status)
	})
	t.Run("When_Master_Is_Not_Appropriate", func(t *testing.T) {
		// Given
		astronaut := astronaut{}
		answer := questions.AstronautAnswer{
			Tall: 120,
			EyeSight: questions.Select{
				Index: 0,
				Value: questions.AstronautCurableEyeSightProblem,
			},
			BachelorDegree: questions.Select{
				Index: 0,
				Value: questions.BiologicalScienceDegree,
			},
			MasterDegree: questions.Select{
				Index: 0,
				Value: questions.OtherDegrees,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "Sorry, There are really strict rules to be selected. Better to prefer another job."

		//Then
		assert.Equal(t, expectedText, check.text)
		assert.False(t, check.Status)
	})
	t.Run("When_Have_Incurable_Eyesight_Problem", func(t *testing.T) {
		// Given
		astronaut := astronaut{}
		answer := questions.AstronautAnswer{
			Tall: 165,
			EyeSight: questions.Select{
				Index: 0,
				Value: questions.AstronautIncurableEyeSightProblem,
			},
			BachelorDegree: questions.Select{
				Index: 0,
				Value: questions.BiologicalScienceDegree,
			},
			MasterDegree: questions.Select{
				Index: 0,
				Value: questions.OtherDegrees,
			},
		}

		byteAnswer, _ := json.Marshal(answer)

		// When
		check, _ := astronaut.Check(byteAnswer)
		expectedText := "Sorry, There are really strict rules to be selected. Better to prefer another job."

		//Then
		assert.Equal(t, expectedText, check.text)
		assert.False(t, check.Status)
	})
}
