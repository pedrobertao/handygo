package strx

import "testing"

func TestToSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic",
			input:    "HelloWorld",
			expected: "hello_world",
		},
		{
			name:     "Complex",
			input:    "HelloWorldMyNameIsHandyGo",
			expected: "hello_world_my_name_is_handy_go",
		},
		{
			name:     "With spaces right",
			input:    "With Space Right     ",
			expected: "with_space_right",
		},
		{
			name:     "With Spaces left",
			input:    "     With Spaces Left",
			expected: "with_spaces_left",
		},
		{
			name:     "Emoji",
			input:    "Go 💙",
			expected: "go_💙",
		},
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnake(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic",
			input:    "Hello World",
			expected: "dlroW olleH",
		},
		{
			name:     "With spaces right",
			input:    "With Space Right     ",
			expected: "     thgiR ecapS htiW",
		},
		{
			name:     "With Spaces left",
			input:    "     With Spaces Left",
			expected: "tfeL secapS htiW     ",
		},
		{
			name:     "Emoji",
			input:    "Go 💙",
			expected: "💙 oG",
		},
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
