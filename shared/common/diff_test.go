package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CreateEntityDiff(t *testing.T) {
	type testData struct {
		IntValue    int
		StringValue string
	}
	{
		diff, _ := CreateEntityDiff(&testData{}, &testData{})
		require.Equal(t, diff, testData{})
	}
	{
		diff, _ := CreateEntityDiff(&testData{}, &testData{IntValue: 1, StringValue: "2"})
		require.Equal(t, diff, testData{IntValue: 1, StringValue: "2"})
	}
	{
		diff, _ := CreateEntityDiff(&testData{}, &testData{IntValue: 1})
		require.Equal(t, diff, testData{IntValue: 1})
	}
	{
		diff, _ := CreateEntityDiff(&testData{IntValue: 2}, &testData{IntValue: 3})
		require.Equal(t, diff, testData{IntValue: 3})
	}
	{
		diff, _ := CreateEntityDiff(&testData{IntValue: 2}, &testData{StringValue: "3"})
		require.Equal(t, diff, testData{StringValue: "3"})
	}
	{
		diff, _ := CreateEntityDiff(&testData{IntValue: 2}, &testData{})
		require.Equal(t, diff, testData{})
	}
}
