package process_transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionWhenItValid(t *testing.T) {

	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := TrasactionDtoOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	repository :=

	usecase := NewProcessTransaction(repository)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
