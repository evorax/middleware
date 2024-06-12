package middle

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Params map[string]string

func MatchPath(routePath, reqPath string) (Params, error) {
	route := strings.Split(routePath, "/")
	req := strings.Split(reqPath, "/")

	if len(route) != len(req) {
		return nil, errors.New("paths have different number of segments")
	}

	params := make(Params)

	for i, segment := range route {
		if segment == "*" {
			continue
		}

		if strings.HasPrefix(segment, ":") {
			name := segment[1:]
			value := req[i]

			name, err := Match(name, value)
			if err != nil {
				Error(err)
			}

			params[name] = value
		} else if segment != req[i] {
			Error(fmt.Sprintf("path segments do not match: '%s' != '%s'", segment, req[i]))
		}
	}

	return params, nil
}

func Match(name, value string) (string, error) {
	if index := strings.Index(name, "["); index != -1 {
		pattern := name[index+1 : len(name)-1]
		name = name[:index]

		match, err := regexp.MatchString(pattern, value)
		if err != nil {
			return "", err
		}

		if !match {
			return "", fmt.Errorf("parameter value '%s' does not match pattern '%s'", value, pattern)
		}
	}
	return name, nil
}
