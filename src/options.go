package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type argument struct {
	Name        string
	Option        string
	Description string
	Default     interface{}
}


func cmdUsage() {
	w := flag.CommandLine.Output()
	fmt.Fprintf(w, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()

}

// Parse a value based on its expected type.
func parseValueType(val string, arg argument) (interface{}, error) {
	switch arg.Default.(type) {
	case *string:
		return val, nil
	case *int:
		conv, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("value '%s' cannot be converted to int", val)
		}
		return conv, nil
	case *bool:
		switch strings.ToUpper(val) {
		case "TRUE", "T":
			return true, nil
		case "FALSE", "F":
			return false, nil
		default:
			return nil, fmt.Errorf("unsupported boolean value '%s'", val)
		}
	default:
		return nil, fmt.Errorf("unsupported type")
	}
}

func parseCmdLine(args []argument) map[string]interface{} {

	parsed := make(map[string]interface{})

	for _, arg := range args {
		switch arg.Default.(type) {
		case string:
			flag.String(arg.Option, arg.Default.(string), arg.Description)
		case int:
			flag.Int(arg.Option, arg.Default.(int), arg.Description)
		case  bool:
			flag.Bool(arg.Option, arg.Default.(bool), arg.Description)
		}
	}

	flag.Usage = cmdUsage
	flag.Parse()

	for _, arg := range args {
		// Get values from flags
		switch arg.Default.(type) {
		case string:
			parsed[arg.Name] = flag.Lookup(arg.Option).Value.String()
		case int:
			val, err := strconv.Atoi(flag.Lookup(arg.Option).Value.String())
			if err != nil {
				fmt.Printf("Error parsing %s to int: %v\n", arg.Name, err)
			} else {
				parsed[arg.Name] = val
			}
		case bool:
			val := flag.Lookup(arg.Option).Value.String()
			parsed[arg.Name] = val == "true"
		}
	}
	return parsed
}

// Parse the arguments from environment variables.
func parseArgs(leader string, args []argument) map[string]interface{} {

	parsed := parseCmdLine(args)

	for _, arg := range args {
		env := os.Getenv(leader + "_" + arg.Name)
		if env != "" {
			val, err := parseValueType(env, arg)
			if err != nil {
				fmt.Printf("Error parsing %s: %v\n", arg.Name, err)
			} else  {
				parsed[arg.Name] = val
				continue
			} 
		}
	}

	return parsed
}


