package format

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
)

// PrintTitle prints a title with lipgloss styling and emojis
func PrintTitle(title string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("57")).
		Bold(true).
		Padding(0, 1).
		Margin(1, 0)

	title = style.Render("✨ " + title + " ✨")
	logrus.Infof("\n\n%s\n", title)
}

// PrintError prints an error message with lipgloss styling and emojis
func PrintError(message string) {
	style := lipgloss.NewStyle().
		Foreground(liposs.Color("9")).
		Bold(true).
		Padding(0, 1).
		Margin(1, 0)

	message = style.Render("❌ " + message + " ❌")
	logrus.Errorf("\n\n%s\n", message)
}

// PrintSuccess prints a success message with lipgloss styling and emojis
func PrintSuccess(message string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true).
		Padding(0, 1).
		Margin(1, 0)

	message = style.Render("✅ " + message + " ✅")
	logrus.Infof("\n\n%s\n", message)
}
