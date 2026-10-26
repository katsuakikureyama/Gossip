// Fixed-profile regexp2 probe for the WordPress block-parser token pattern.
//
// This deliberately accepts only a pattern body, one raw subject, and a
// normalized byte offset. It is comparison evidence plumbing, not a general
// regexp runner or a PHP compatibility layer.
package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/dlclark/regexp2/v2"
)

const (
	inputLimit = 4096
	stackSlots = 100000
	timeoutMS  = 100
)

type probeGroup struct {
	Number       int    `json:"number"`
	Name         string `json:"name"`
	Participated bool   `json:"participated"`
	RuneIndex    int    `json:"rune_index"`
	RuneLength   int    `json:"rune_length"`
	ByteStart    int    `json:"byte_start"`
	ByteLength   int    `json:"byte_length"`
	Base64       string `json:"base64"`
}

type probeResult struct {
	Status     string       `json:"status"`
	Error      string       `json:"error"`
	ErrorClass string       `json:"error_class"`
	Options    int          `json:"options"`
	StackSlots int          `json:"stack_slots"`
	TimeoutMS  int          `json:"timeout_ms"`
	Groups     []probeGroup `json:"groups"`
	GoVersion  string       `json:"go_version"`
}

func main() {
	os.Exit(run())
}

func run() int {
	defer regexp2.StopTimeoutClock()

	if len(os.Args) != 4 {
		return usage("expected: parser-regexp2-probe <pattern-body-file> <subject-file> <normalized-byte-offset>")
	}

	pattern, err := readBoundedRegularFile(os.Args[1])
	if err != nil {
		return usage(err.Error())
	}
	subject, err := readBoundedRegularFile(os.Args[2])
	if err != nil {
		return usage(err.Error())
	}
	offset, err := parseOffset(os.Args[3], len(subject))
	if err != nil {
		return usage(err.Error())
	}

	options := regexp2.RE2 | regexp2.Singleline
	result := probeResult{
		Options:    int(options),
		StackSlots: stackSlots,
		TimeoutMS:  timeoutMS,
		Groups:     make([]probeGroup, 0),
		GoVersion:  runtime.Version(),
	}

	re, err := regexp2.Compile(string(pattern), options, regexp2.OptionMaxBacktrackingStackSize(stackSlots))
	if err != nil {
		result.Status = "error"
		result.Error = err.Error()
		result.ErrorClass = "compile"
		return writeResult(result)
	}
	re.MatchTimeout = timeoutMS * time.Millisecond

	match, err := re.FindStringMatchStartingAt(string(subject), offset)
	if err != nil {
		result.Status = "error"
		result.Error = err.Error()
		result.ErrorClass = matchErrorClass(err, subject, offset)
		return writeResult(result)
	}
	if match == nil {
		result.Status = "no-match"
		return writeResult(result)
	}

	result.Status = "match"
	numbers := re.GetGroupNumbers()
	names := re.GetGroupNames()
	for index, number := range numbers {
		name := ""
		if index < len(names) {
			name = names[index]
		}
		group := match.GroupByNumber(number)
		if group == nil {
			return usage(fmt.Sprintf("regexp2 returned no group for declared group number %d", number))
		}

		observation := probeGroup{
			Number:       number,
			Name:         name,
			Participated: len(group.Captures) > 0,
			RuneIndex:    group.RuneIndex,
			RuneLength:   group.RuneLength,
			ByteStart:    -1,
		}
		if observation.Participated {
			observation.ByteStart, observation.ByteLength = group.ByteRange()
			observation.Base64 = base64.StdEncoding.EncodeToString([]byte(group.String()))
		}
		result.Groups = append(result.Groups, observation)
	}
	return writeResult(result)
}

func readBoundedRegularFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open input file %q: %w", path, err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("cannot stat input file %q: %w", path, err)
	}
	if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("input must be a regular file: %q", path)
	}
	if info.Size() > inputLimit {
		return nil, fmt.Errorf("input exceeds %d bytes: %q", inputLimit, path)
	}

	data, err := io.ReadAll(io.LimitReader(file, inputLimit+1))
	if err != nil {
		return nil, fmt.Errorf("cannot read input file %q: %w", path, err)
	}
	if len(data) > inputLimit {
		return nil, fmt.Errorf("input changed or exceeds %d bytes: %q", inputLimit, path)
	}
	return data, nil
}

func parseOffset(value string, subjectLength int) (int, error) {
	if value == "" {
		return 0, errors.New("normalized byte offset must be a non-negative decimal integer")
	}
	for _, character := range value {
		if character < '0' || character > '9' {
			return 0, errors.New("normalized byte offset must be a non-negative decimal integer")
		}
	}
	offset, err := strconv.Atoi(value)
	if err != nil || offset > subjectLength {
		return 0, fmt.Errorf("normalized byte offset must be within the subject: %q", value)
	}
	return offset, nil
}

func matchErrorClass(err error, subject []byte, offset int) string {
	if errors.Is(err, regexp2.ErrBacktrackingStackLimit) {
		return "backtracking-stack-limit"
	}
	if insideRuneOffset(subject, offset) {
		return "inside-rune-offset"
	}
	return "match-error"
}

func insideRuneOffset(subject []byte, offset int) bool {
	if offset <= 0 || offset >= len(subject) {
		return false
	}
	for index := range string(subject) {
		if index == offset {
			return false
		}
		if index > offset {
			return true
		}
	}
	return false
}

func writeResult(result probeResult) int {
	encoded, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parser-regexp2-probe: cannot encode result: %v\n", err)
		return 1
	}
	if _, err := os.Stdout.Write(append(encoded, '\n')); err != nil {
		fmt.Fprintf(os.Stderr, "parser-regexp2-probe: cannot write result: %v\n", err)
		return 1
	}
	return 0
}

func usage(message string) int {
	fmt.Fprintf(os.Stderr, "parser-regexp2-probe: %s\n", message)
	return 1
}
