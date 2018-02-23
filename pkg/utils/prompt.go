package utils

// Prompt represents different types of prompts
type Prompt string

const (
	// PromptCaret is a visually appealing prompt
	PromptCaret Prompt = "> "
	// PromptNone is a test friendly prompt
	PromptNone Prompt = ""
)
