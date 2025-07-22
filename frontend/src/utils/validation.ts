// Frontend validation utility for entity creation and editing

export interface ValidationResult {
  isValid: boolean;
  error?: string;
  field?: string;
}

export interface ValidationConfig {
  minLength?: number;
  maxLength?: number;
  required?: boolean;
  allowEmpty?: boolean;
}

export class EntityValidator {
  private static readonly DEFAULT_CONFIG = {
    minLength: 1,
    maxLength: 200,
    required: true,
    allowEmpty: false,
  };

  private static readonly DESCRIPTION_CONFIG = {
    minLength: 1,
    maxLength: 2000,
    required: true,
    allowEmpty: false,
  };

  /**
   * Sanitizes a string by trimming whitespace
   */
  static sanitize(value: string): string {
    return value.trim();
  }

  /**
   * Validates a string field with given configuration
   */
  static validateField(
    value: string,
    fieldName: string,
    entityType: string,
    config: ValidationConfig = EntityValidator.DEFAULT_CONFIG
  ): ValidationResult {
    const sanitized = EntityValidator.sanitize(value);
    const { minLength, maxLength, required, allowEmpty } = { ...EntityValidator.DEFAULT_CONFIG, ...config };

    // Check if field is required
    if (required && sanitized === '') {
      return {
        isValid: false,
        error: `${entityType} ${fieldName} is required and cannot be empty`,
        field: fieldName,
      };
    }

    // If empty and allowed, return valid
    if (sanitized === '' && allowEmpty) {
      return { isValid: true };
    }

    // Check minimum length
    if (sanitized.length < minLength) {
      return {
        isValid: false,
        error: `${entityType} ${fieldName} must be at least ${minLength} character${minLength > 1 ? 's' : ''} long`,
        field: fieldName,
      };
    }

    // Check maximum length
    if (sanitized.length > maxLength) {
      return {
        isValid: false,
        error: `${entityType} ${fieldName} cannot exceed ${maxLength} characters`,
        field: fieldName,
      };
    }

    return { isValid: true };
  }

  /**
   * Validates a character name
   */
  static validateCharacterName(name: string): ValidationResult {
    return EntityValidator.validateField(name, 'name', 'Character', EntityValidator.DEFAULT_CONFIG);
  }

  /**
   * Validates a character description
   */
  static validateCharacterDescription(description: string): ValidationResult {
    return EntityValidator.validateField(description, 'description', 'Character', EntityValidator.DESCRIPTION_CONFIG);
  }

  /**
   * Validates a location name
   */
  static validateLocationName(name: string): ValidationResult {
    return EntityValidator.validateField(name, 'name', 'Location', EntityValidator.DEFAULT_CONFIG);
  }

  /**
   * Validates a location description
   */
  static validateLocationDescription(description: string): ValidationResult {
    return EntityValidator.validateField(description, 'description', 'Location', EntityValidator.DESCRIPTION_CONFIG);
  }

  /**
   * Validates a codex entry title
   */
  static validateCodexTitle(title: string): ValidationResult {
    return EntityValidator.validateField(title, 'title', 'Codex entry', EntityValidator.DEFAULT_CONFIG);
  }

  /**
   * Validates a codex entry content
   */
  static validateCodexContent(content: string): ValidationResult {
    return EntityValidator.validateField(content, 'content', 'Codex entry', EntityValidator.DESCRIPTION_CONFIG);
  }

  /**
   * Validates a complete character object
   */
  static validateCharacter(name: string, description: string): ValidationResult {
    const nameValidation = EntityValidator.validateCharacterName(name);
    if (!nameValidation.isValid) {
      return nameValidation;
    }

    const descValidation = EntityValidator.validateCharacterDescription(description);
    if (!descValidation.isValid) {
      return descValidation;
    }

    return { isValid: true };
  }

  /**
   * Validates a complete location object
   */
  static validateLocation(name: string, description: string): ValidationResult {
    const nameValidation = EntityValidator.validateLocationName(name);
    if (!nameValidation.isValid) {
      return nameValidation;
    }

    const descValidation = EntityValidator.validateLocationDescription(description);
    if (!descValidation.isValid) {
      return descValidation;
    }

    return { isValid: true };
  }

  /**
   * Validates a complete codex entry object
   */
  static validateCodexEntry(title: string, content: string): ValidationResult {
    const titleValidation = EntityValidator.validateCodexTitle(title);
    if (!titleValidation.isValid) {
      return titleValidation;
    }

    const contentValidation = EntityValidator.validateCodexContent(content);
    if (!contentValidation.isValid) {
      return contentValidation;
    }

    return { isValid: true };
  }

  /**
   * Validates and sanitizes multiple fields at once
   */
  static validateAndSanitizeFields(
    fields: Array<{
      value: string;
      name: string;
      entityType: string;
      config?: ValidationConfig;
    }>
  ): { isValid: boolean; errors: string[]; sanitizedValues: string[] } {
    const errors: string[] = [];
    const sanitizedValues: string[] = [];

    for (const field of fields) {
      const sanitized = EntityValidator.sanitize(field.value);
      sanitizedValues.push(sanitized);

      const validation = EntityValidator.validateField(
        field.value,
        field.name,
        field.entityType,
        field.config
      );

      if (!validation.isValid && validation.error) {
        errors.push(validation.error);
      }
    }

    return {
      isValid: errors.length === 0,
      errors,
      sanitizedValues,
    };
  }
}

// Export validation constants for external use
export const VALIDATION_CONSTANTS = {
  NAME_MIN_LENGTH: 1,
  NAME_MAX_LENGTH: 200,
  DESCRIPTION_MIN_LENGTH: 1,
  DESCRIPTION_MAX_LENGTH: 2000,
} as const;

// Helper function for form validation
export function createFormValidator(entityType: string) {
  return {
    validateName: (name: string) => EntityValidator.validateField(name, 'name', entityType),
    validateDescription: (description: string) => 
      EntityValidator.validateField(description, 'description', entityType, EntityValidator.DESCRIPTION_CONFIG),
    validateComplete: (name: string, description: string) => {
      const nameResult = EntityValidator.validateField(name, 'name', entityType);
      if (!nameResult.isValid) return nameResult;
      
      const descResult = EntityValidator.validateField(description, 'description', entityType, EntityValidator.DESCRIPTION_CONFIG);
      return descResult;
    },
  };
}
