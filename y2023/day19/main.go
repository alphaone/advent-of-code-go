package day19

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils"
)

type Workflow struct {
	Name   string
	Stages []Stage
}

type Stage struct {
	Condition   Condition
	CatchAll    bool
	Destination string
}

type Condition struct {
	Field    string
	Operator rune
	Value    int
}

type Part struct{ X, M, A, S int }

func parseWorkflow(s string) Workflow {
	wf := Workflow{}
	re := regexp.MustCompile(`(\w+){(.*)}`)
	res := re.FindStringSubmatch(s)
	wf.Name = res[1]

	for _, stage := range strings.Split(res[2], ",") {
		reStage := regexp.MustCompile(`(\w+)([<>])(\d+):(\w+)`)
		resStage := reStage.FindStringSubmatch(stage)
		if len(resStage) == 0 {
			wf.Stages = append(wf.Stages, Stage{CatchAll: true, Destination: stage})
			continue
		}
		value, _ := strconv.Atoi(resStage[3])
		wf.Stages = append(wf.Stages, Stage{
			Condition: Condition{
				Field:    resStage[1],
				Operator: rune(resStage[2][0]),
				Value:    value,
			},
			Destination: resStage[4],
		})
	}
	return wf
}

func parseWorkflows(ss []string) map[string]Workflow {
	wfs := map[string]Workflow{}
	for _, s := range ss {
		wf := parseWorkflow(s)
		wfs[wf.Name] = wf
	}
	return wfs
}

func parsePart(s string) Part {
	part := Part{}
	re := regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
	res := re.FindStringSubmatch(s)
	part.X, _ = strconv.Atoi(res[1])
	part.M, _ = strconv.Atoi(res[2])
	part.A, _ = strconv.Atoi(res[3])
	part.S, _ = strconv.Atoi(res[4])
	return part
}

func parseParts(ss []string) []Part {
	var parts []Part
	for _, s := range ss {
		parts = append(parts, parsePart(s))
	}
	return parts
}

func (p Part) Apply(wf Workflow) string {
	for _, stage := range wf.Stages {
		if stage.CatchAll {
			return stage.Destination
		}
		if stage.Condition.Apply(p) {
			return stage.Destination
		}
	}
	panic("no catch-all stage found")
}

func (c Condition) Apply(p Part) bool {
	switch c.Field {
	case "x":
		return applyOperator(c.Value, p.X, c.Operator)
	case "m":
		return applyOperator(c.Value, p.M, c.Operator)
	case "a":
		return applyOperator(c.Value, p.A, c.Operator)
	case "s":
		return applyOperator(c.Value, p.S, c.Operator)
	default:
		panic("unknown field")
	}
}

func applyOperator(condVal, partVal int, op rune) bool {
	switch op {
	case '<':
		return partVal < condVal
	case '>':
		return partVal > condVal
	default:
		panic("unknown operator")
	}
}

func (p Part) ApplyAll(wfs map[string]Workflow) []string {
	var results []string
	wf := wfs["in"]
	for {
		wfName := p.Apply(wf)
		results = append(results, wfName)
		if wfName == "R" || wfName == "A" {
			break
		}
		wf = wfs[wfName]
	}
	return results
}

func solvePartA(input []string) int {
	sections := utils.SplitIntoSection(input)
	wfs := parseWorkflows(sections[0])
	parts := parseParts(sections[1])

	accepted := []Part{}
	for _, part := range parts {
		res := part.ApplyAll(wfs)
		if res[len(res)-1] == "A" {
			accepted = append(accepted, part)
		}
	}

	result := 0
	for _, part := range accepted {
		result += part.X + part.M + part.A + part.S
	}
	return result
}
