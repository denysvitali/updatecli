package format

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
)

// PrintError prints an error message with lipgloss styling and emojis
func PrintError(message string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Bold(true).
		Padding(0, 1).
		Margin(1, 0)

	message = style.Render("❌ " + message + " ❌")
	logrus.Errorf("\n\n%s\n", message)
}
