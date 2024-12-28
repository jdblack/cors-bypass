package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type argument struct {
	Short       string
	Long        string
	Optional    bool
	Description string
	Type        interface{}
	Default     interface{}
}

type parsedArg struct {
	Name string
	Val  interface{}
}

// Parse a value based on its expected type.
func parseValueType(val interface{}, argtype interface{}) (interface{}, error) {
	switch argtype.(type) {
	case *string:
		return val, nil
	case *int:
		if strVal, ok := val.(string); ok {
			conv, err := strconv.Atoi(strVal)
			if err != nil {
				return nil, fmt.Errorf("value '%s' cannot be converted to int", strVal)
			}
			return conv, nil
		}
		return nil, fmt.Errorf("value is not a string for int type")
	case *bool:
		if strVal, ok := val.(string); ok {
			switch strings.ToUpper(strVal) {
			case "TRUE", "T":
				return true, nil
			case "FALSE", "F":
				return false, nil
			default:
				return nil, fmt.Errorf("unsupported boolean value '%s'", strVal)
			}
		}
		return nil, fmt.Errorf("value is not a string for bool type")
	default:
		return nil, fmt.Errorf("unsupported type")
	}
}

// Parse the arguments from environment variables.
func parseArgs(leader string, args []argument) []parsedArg {
	var parsed []parsedArg

	for _, arg := range args {
		env := os.Getenv(leader + "_" + arg.Long)
		if env != "" {
			val, err := parseValueType(env, arg.Type)
			if err == nil {
				parsed = append(parsed, parsedArg{Name: arg.Long, Val: val})
			} else {
				fmt.Printf("Error parsing %s: %v\n", arg.Long, err)
			}
			continue
		}

		if arg.Default != nil {
			val, err := parseValueType(arg.Default, arg.Type)
			if err == nil {
				parsed = append(parsed, parsedArg{Name: arg.Long, Val: val})
			} else {
				fmt.Printf("Error parsing default value for %s: %v\n", arg.Long, err)
			}
		}
	}

	return parsed
}

func main() {
	// Sample argument definitions with correct default types
	args := []argument{
		{Short: "h", Long: "help", Optional: true, Description: "Display help", Type: new(bool), Default: "false"},
		{Short: "v", Long: "version", Optional: false, Description: "Version number", Type: new(string), Default: "1.0.0"},
		{Short: "n", Long: "number", Optional: false, Description: "An integer number", Type: new(int), Default: "10"},
	}

	// Set a sample environment variable for testing
	os.Setenv("LEADER_VERSION", "2.0.0")
	os.Setenv("LEADER_NUMBER", "42") // Add this for testing number parsing

	// Parse the arguments
	parsedArgs := parseArgs("LEADER", args)

	// Print the parsed arguments
	for _, p := range parsedArgs {
		fmt.Printf("Name: %s, Val: %v\n", p.Name, p.Val)
	}
}


