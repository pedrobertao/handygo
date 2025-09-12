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
			name:     "Already Snake",
			input:    "hello_world",
			expected: "hello_world",
		},
		{
			name:     "Edge Case 1",
			input:    "Hello_World_",
			expected: "hello_world_",
		},
		{
			name:     "Edge Case 2",
			input:    "_Hello_World",
			expected: "_hello_world",
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
		{
			name:     "Empty Space 1",
			input:    " ",
			expected: " ",
		},
		{
			name:     "Empty Space 2",
			input:    "    ",
			expected: "    ",
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

func TestToKebab(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic camelCase",
			input:    "HelloWorld",
			expected: "hello-world",
		},
		{
			name:     "Complex camelCase",
			input:    "HelloWorldMyNameIsHandyGo",
			expected: "hello-world-my-name-is-handy-go",
		},
		{
			name:     "With spaces",
			input:    "Hello World",
			expected: "hello-world",
		},
		{
			name:     "Multiple spaces",
			input:    "Hello   World   Test",
			expected: "hello-world-test",
		},
		{
			name:     "Already kebab",
			input:    "hello-world",
			expected: "hello-world",
		},
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "Empty Space 1",
			input:    " ",
			expected: " ",
		},
		{
			name:     "Empty Space 2",
			input:    "    ",
			expected: "    ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToKebab(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToScreamingSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic camelCase",
			input:    "HelloWorld",
			expected: "HELLO_WORLD",
		},
		{
			name:     "Complex camelCase",
			input:    "HelloWorldMyNameIsHandyGo",
			expected: "HELLO_WORLD_MY_NAME_IS_HANDY_GO",
		},
		{
			name:     "With spaces right",
			input:    "With Spaces Right     ",
			expected: "WITH_SPACES_RIGHT",
		},
		{
			name:     "With Spaces left",
			input:    "     With Spaces Left",
			expected: "WITH_SPACES_LEFT",
		},
		{
			name:     "Already Screaming SnakeCase",
			input:    "HELLO_WORLD",
			expected: "HELLO_WORLD",
		},
		{
			name:     "Edge Case 1",
			input:    "Hello_World_",
			expected: "HELLO_WORLD_",
		},
		{
			name:     "Edge Case 2",
			input:    "_Hello_World",
			expected: "_HELLO_WORLD",
		},
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "Empty Space 1",
			input:    " ",
			expected: " ",
		},
		{
			name:     "Empty Space 2",
			input:    "    ",
			expected: "    ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToScreamingSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Snake case",
			input:    "hello_world",
			expected: "helloWorld",
		},
		{
			name:     "Complex snake case",
			input:    "hello_world_my_name_is_handy_go",
			expected: "helloWorldMyNameIsHandyGo",
		},
		{
			name:     "Kebab case",
			input:    "hello-world",
			expected: "helloWorld",
		},
		{
			name:     "Complex kebab case",
			input:    "hello-world-my-name",
			expected: "helloWorldMyName",
		},
		{
			name:     "With spaces",
			input:    "Hello World",
			expected: "helloWorld",
		},
		{
			name:     "Multiple spaces",
			input:    "Hello   World   Test",
			expected: "helloWorldTest",
		},
		{
			name:     "Already camelCase",
			input:    "HelloWorld",
			expected: "helloWorld",
		},
		{
			name:     "With spaces and trim",
			input:    "  hello world  ",
			expected: "helloWorld",
		},
		{
			name:     "Empty underscores",
			input:    "hello__world",
			expected: "helloWorld",
		},
		{
			name:     "Empty dashes",
			input:    "hello--world",
			expected: "helloWorld",
		},
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "Only spaces",
			input:    "   ",
			expected: "",
		},
		{
			name:     "Single word",
			input:    "Hello",
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToCamel(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
