package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type arguments struct {
	users       int
	links       []string
	intervalMin int
	intervalMax int
	throttle    bool
	duration    int
}

type problem struct {
	desc string
}

func splitArgs(args []string, fn func(arg1 string, arg2 string)) {
	for _, arg := range args {
		split := strings.Split(arg, "=")

		if len(split) != 2 {
			fn(split[0], "")

			continue
		}

		fn(split[0], split[1])
	}
}

func createArgsMap(args []string) map[string]string {
	argMap := make(map[string]string)
	splitArgs(args, func(name string, value string) {
		argMap[name] = value
	})

	return argMap
}

func validateValid(args map[string]string, valid []string) []problem {
	invalid := make([]string, 0)
	for k, _ := range args {
		found := false
		for _, v := range valid {
			if k == v {
				found = true
				break
			}
		}

		if !found {
			invalid = append(invalid, k)
		}
	}

	problems := make([]problem, 0)
	for _, v := range invalid {
		problems = append(problems, problem{fmt.Sprintf("Invalid argument %s. Run --help to view valid arguments.", v)})
	}

	return problems
}

func validateRequired(args map[string]string, required []string) problem {
	for k, _ := range args {
		found := false
		for _, v := range required {
			if k == v {
				found = true
				break
			}
		}

		if found {
			return problem{}
		}
	}

	return problem{desc: "--links argument is required."}
}

func processLinks(links string) []string {
	return strings.Split(links, ",")
}

func processThrottle(argsMap map[string]string) bool {
	_, ok := argsMap["--throttle"]

	return ok
}

func processUsers(argMap map[string]string) (int, problem) {
	if _, ok := argMap["--users"]; !ok {
		return 50, problem{}
	}

	value := argMap["--users"]

	j, err := strconv.Atoi(value)

	if err != nil {
		return 0, problem{"Unable to convert --users value to integer. --users value must be an integer"}
	}

	return j, problem{}
}

func processDuration(argMap map[string]string) (int, problem) {
	if _, ok := argMap["--duration"]; !ok {
		return 60, problem{}
	}

	value := argMap["--duration"]

	j, err := strconv.Atoi(value)

	if err != nil {
		return 0, problem{"Unable to convert --duration value to integer. --duration value must be an integer"}
	}

	return j, problem{}
}

func processInterval(argMap map[string]string) (int, int, problem) {
	if _, ok := argMap["--interval"]; !ok {
		return 3, 15, problem{}
	}

	value := argMap["--interval"]

	s := strings.Split(value, "-")

	if len(s) == 2 {
		m, err := strconv.Atoi(s[0])

		if err != nil {
			return 0, 0, problem{"Invalid --interval argument. --interval argument can be either a single number (for max) or a range {min}-{max}. For example, 3 for maximum value or 3-15 for range. min cannot be less than 3."}
		}

		if m < 3 {
			return 0, 0, problem{"Invalid --interval argument. min cannot be less that 3"}
		}

		n, err := strconv.Atoi(s[1])

		if err != nil {
			return 0, 0, problem{"Invalid --interval argument. --interval argument can be either a single number (for max) or a range {min}-{max}. For example, 3 for maximum value or 3-15 for range. min cannot be less than 3."}
		}

		return m, n, problem{}
	}

	if len(s) == 1 {
		m, err := strconv.Atoi(s[0])

		if err != nil {
			return 0, 0, problem{"Invalid --interval argument. --interval argument can be either a single number (for max) or a range {min}-{max}. For example, 3 for maximum value or 3-15 for range. min cannot be less than 3."}
		}

		return m, 15, problem{}
	}

	return 0, 0, problem{}
}

func processProblems(problems []problem) error {
	found := false
	str := `
Invalid arguments. The following problems were found:

`
	for _, problem := range problems {
		if problem.desc != "" {
			found = true

			str += fmt.Sprintf("->  %s\n", problem.desc)
		}
	}

	if found {
		fmt.Println(str)
		fmt.Println("Run rosie --help for more information.")
		fmt.Println("")

		return errors.New("Invalid arguments")
	}

	return nil
}

func newArgs(args []string) (arguments, bool) {
	argMap := createArgsMap(args)

	if _, ok := argMap["--help"]; ok {

	}

	problems := make([]problem, 0)
	problems = append(problems, validateValid(argMap, []string{"--links", "--users", "--duration", "--throttle", "--interval"})...)
	problems = append(problems, validateRequired(argMap, []string{"--links"}))

	users, problem := processUsers(argMap)
	problems = append(problems, problem)
	duration, problem := processDuration(argMap)
	problems = append(problems, problem)
	min, max, problem := processInterval(argMap)
	problems = append(problems, problem)
	links := processLinks(argMap["--links"])

	if err := processProblems(problems); err != nil {
		return arguments{}, false
	}

	return arguments{
		users:       users,
		links:       links,
		intervalMin: min,
		intervalMax: max,
		throttle:    processThrottle(argMap),
		duration:    duration,
	}, true

	/*	_, ok := argMap["--links"]

		if !ok {
			return arguments{}, errors.New("--links argument is required. It must a comma separated list of valid http URLs. For example, --links=https://google.com,https://facebook.com")
		}

		for k, v := range argMap {
			if k == "--links" {
				a.links = strings.Split(v, ",")
			}

			if k == "--users" {
				j, err := strconv.Atoi(v)

				if err != nil {
					return arguments{}, errors.New("Unable to convert --users value to integer. --users value must be an integer")
				}

				a.users = j
			}

			if k == "--interval" {
				s := strings.Split(v, "-")

				if len(s) == 2 {
					m, err := strconv.Atoi(s[0])

					if err != nil {
						return arguments{}, errors.New("Invalid --interval argument. --interval argument can be either a single number (for max) or a range {min}-{max}. For example, 3 for maximum value or 3-15 for range. min cannot be less than 3.")
					}

					if m < 3 {
						return arguments{}, errors.New("Invalid --interval argument. min cannot be less that 3")
					}

					n, err := strconv.Atoi(s[0])

					if err != nil {
						return arguments{}, errors.New("Invalid --interval argument. --interval argument can be either a single number (for max) or a range {min}-{max}. For example, 3 for maximum value or 3-15 for range. min cannot be less than 3.")
					}

					a.intervalMin = m
					a.intervalMax = n
				}

				if len(s) == 1 {
					m, err := strconv.Atoi(s[0])

					if err != nil {
						return arguments{}, errors.New("Invalid --interval argument. --interval argument can be either a single number (for max) or a range {min}-{max}. For example, 3 for maximum value or 3-15 for range. min cannot be less than 3.")
					}

					a.intervalMax = m
				}
			}

			if k == "--throttle" && v == "true" {
				a.throttle = true
			}

			if k == "--duration" {
				m, err := strconv.Atoi(v)

				if err != nil {
					return arguments{}, errors.New("Unable to convert --duration value to integer. --duration value must be an integer")
				}

				a.duration = m
			}
		}
	*/
}
