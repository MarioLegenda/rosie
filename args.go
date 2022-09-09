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
}

func newArgs(args []string) (arguments, error) {
	if len(args) == 0 {
		return arguments{}, errors.New("Argument list is empty. run --help to see the list of arguments.")
	}

	argMap := make(map[string]string)
	for _, arg := range args {
		split := strings.Split(arg, "=")

		if len(split) != 2 {
			return arguments{}, errors.New(fmt.Sprintf("Invalid argument %s given. Arguments must follow --arg=value convention.", split[0]))
		}

		argMap[split[0]] = split[1]
	}

	a := arguments{
		users:       0,
		links:       nil,
		intervalMin: 0,
		intervalMax: 0,
	}

	_, ok := argMap["--links"]

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
				return arguments{}, errors.New("Unable to converts --users value to integer. --users value must be an integer")
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
	}

	return a, nil
}
