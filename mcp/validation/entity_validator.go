package validation

import (
	"fmt"
	"strings"
)

// Constants for validation
const (
	MinNameLength        = 1
	MinDescriptionLength = 1
	MaxNameLength        = 200
	MaxDescriptionLength = 2000
)

// ValidationError represents a validation error with field and message
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// EntityValidator provides validation for common entity fields
type EntityValidator struct{}

// NewEntityValidator creates a new entity validator
func NewEntityValidator() *EntityValidator {
	return &EntityValidator{}
}

// ValidateName validates entity name field
func (v *EntityValidator) ValidateName(name string, entityType string) error {
	name = strings.TrimSpace(name)
	
	if name == "" {
		return ValidationError{
			Field:   "name",
			Message: fmt.Sprintf("%s name cannot be empty", entityType),
		}
	}
	
	if len(name) < MinNameLength {
		return ValidationError{
			Field:   "name",
			Message: fmt.Sprintf("%s name must be at least %d character(s)", entityType, MinNameLength),
		}
	}
	
	if len(name) > MaxNameLength {
		return ValidationError{
			Field:   "name", 
			Message: fmt.Sprintf("%s name cannot exceed %d characters", entityType, MaxNameLength),
		}
	}
	
	return nil
}

// ValidateDescription validates entity description field
func (v *EntityValidator) ValidateDescription(description string, entityType string, required bool) error {
	description = strings.TrimSpace(description)
	
	if required && description == "" {
		return ValidationError{
			Field:   "description",
			Message: fmt.Sprintf("%s description cannot be empty", entityType),
		}
	}
	
	if description != "" && len(description) < MinDescriptionLength {
		return ValidationError{
			Field:   "description",
			Message: fmt.Sprintf("%s description must be at least %d character(s)", entityType, MinDescriptionLength),
		}
	}
	
	if len(description) > MaxDescriptionLength {
		return ValidationError{
			Field:   "description",
			Message: fmt.Sprintf("%s description cannot exceed %d characters", entityType, MaxDescriptionLength),
		}
	}
	
	return nil
}

// ValidateTitle validates entity title field (for codex entries, etc.)
func (v *EntityValidator) ValidateTitle(title string, entityType string) error {
	title = strings.TrimSpace(title)
	
	if title == "" {
		return ValidationError{
			Field:   "title",
			Message: fmt.Sprintf("%s title cannot be empty", entityType),
		}
	}
	
	if len(title) < MinNameLength {
		return ValidationError{
			Field:   "title",
			Message: fmt.Sprintf("%s title must be at least %d character(s)", entityType, MinNameLength),
		}
	}
	
	if len(title) > MaxNameLength {
		return ValidationError{
			Field:   "title",
			Message: fmt.Sprintf("%s title cannot exceed %d characters", entityType, MaxNameLength),
		}
	}
	
	return nil
}

// ValidateContent validates entity content field (for codex entries, etc.)
func (v *EntityValidator) ValidateContent(content string, entityType string) error {
	content = strings.TrimSpace(content)
	
	if content == "" {
		return ValidationError{
			Field:   "content",
			Message: fmt.Sprintf("%s content cannot be empty", entityType),
		}
	}
	
	if len(content) < MinDescriptionLength {
		return ValidationError{
			Field:   "content",
			Message: fmt.Sprintf("%s content must be at least %d character(s)", entityType, MinDescriptionLength),
		}
	}
	
	return nil
}

// ValidateCharacter validates a complete character
func (v *EntityValidator) ValidateCharacter(name, description string) error {
	if err := v.ValidateName(name, "Character"); err != nil {
		return err
	}
	
	if err := v.ValidateDescription(description, "Character", true); err != nil {
		return err
	}
	
	return nil
}

// ValidateLocation validates a complete location
func (v *EntityValidator) ValidateLocation(name, description string) error {
	if err := v.ValidateName(name, "Location"); err != nil {
		return err
	}
	
	if err := v.ValidateDescription(description, "Location", true); err != nil {
		return err
	}
	
	return nil
}

// ValidateCodexEntry validates a complete codex entry
func (v *EntityValidator) ValidateCodexEntry(title, content string) error {
	if err := v.ValidateTitle(title, "Codex entry"); err != nil {
		return err
	}
	
	if err := v.ValidateContent(content, "Codex entry"); err != nil {
		return err
	}
	
	return nil
}

// SanitizeString trims whitespace and returns the cleaned string
func (v *EntityValidator) SanitizeString(input string) string {
	return strings.TrimSpace(input)
}

// ValidateAndSanitize validates and sanitizes a string field
func (v *EntityValidator) ValidateAndSanitize(input, fieldType, entityType string, required bool) (string, error) {
	sanitized := v.SanitizeString(input)
	
	switch fieldType {
	case "name":
		if err := v.ValidateName(sanitized, entityType); err != nil {
			return "", err
		}
	case "description":
		if err := v.ValidateDescription(sanitized, entityType, required); err != nil {
			return "", err
		}
	case "title":
		if err := v.ValidateTitle(sanitized, entityType); err != nil {
			return "", err
		}
	case "content":
		if err := v.ValidateContent(sanitized, entityType); err != nil {
			return "", err
		}
	default:
		if required && sanitized == "" {
			return "", ValidationError{
				Field:   fieldType,
				Message: fmt.Sprintf("%s %s cannot be empty", entityType, fieldType),
			}
		}
	}
	
	return sanitized, nil
}
