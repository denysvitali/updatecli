package format

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPrintTitle(t *testing.T) {
	// Mock logrus output
	var logOutput string
	logrus.SetOutput(&logOutput)

	PrintTitle("Test Title")

	expectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("57")).
		Bold(true).
		Padding(0, 1).
		Margin(1, 0)

	expectedOutput := expectedStyle.Render("✨ Test Title ✨")
	assert.Contains(t, logOutput, expectedOutput)
}

func TestPrintError(t *testing.T) {
	// Mock logrus output
	var logOutput string
	logrus.SetOutput(&logOutput)

	PrintError("Test Error")

	expectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Bold(true).
		Padding(0, 1).
		Margin(1, 0)

	expectedOutput := expectedStyle.Render("❌ Test Error ❌")
	assert.Contains(t, logOutput, expectedOutput)
}

func TestPrintSuccess(t *testing.T) {
	// Mock logrus output
	var logOutput string
	logrus.SetOutput(&logOutput)

	PrintSuccess("Test Success")

	expectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true).
		Padding(0, 1).
		Margin(1, 0)

	expectedOutput := expectedStyle.Render("✅ Test Success ✅")
	assert.Contains(t, logOutput, expectedOutput)
}
