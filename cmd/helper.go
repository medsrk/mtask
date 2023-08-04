package cmd

import (
	"meditasker/domain"
	"strings"
	"time"
)

func parseTaskProperties(taskString string) domain.Task {
	var task domain.Task
	firstColonIndex := strings.Index(taskString, ":")

	// If no colon was found, it means that the whole string is the description
	if firstColonIndex == -1 {
		task.Description = taskString
		return task
	}

	// Find the beginning of the word before the colon
	descriptionEndIndex := strings.LastIndex(taskString[:firstColonIndex], " ")

	// If no space was found, the entire string is the description
	if descriptionEndIndex == -1 {
		task.Description = taskString
		return task
	}

	// The description is everything before the first property
	task.Description = strings.TrimSpace(taskString[:descriptionEndIndex])

	// The rest is a series of properties
	properties := taskString[descriptionEndIndex:]

	// Split the properties by space, each property is in the format "key:value"
	for _, property := range strings.Split(properties, " ") {
		// Split the property into key and value
		propertyParts := strings.SplitN(property, ":", 2)

		// Assign the value to the corresponding task field
		switch propertyParts[0] {
		case "project":
			task.Project = propertyParts[1]
		case "due":
			dueTime, _ := time.Parse(time.RFC3339, propertyParts[1])
			task.Due = dueTime
		}
	}

	return task
}
