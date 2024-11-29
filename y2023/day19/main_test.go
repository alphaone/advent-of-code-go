package day19

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"px{a<2006:qkq,m>2090:A,rfg}",
	"pv{a>1716:R,A}",
	"lnx{m>1548:A,A}",
	"rfg{s<537:gd,x>2440:R,A}",
	"qs{s>3448:A,lnx}",
	"qkq{x<1416:A,crn}",
	"crn{x>2662:A,R}",
	"in{s<1351:px,qqz}",
	"qqz{s>2770:qs,m<1801:hdj,R}",
	"gd{a>3333:R,R}",
	"hdj{m>838:A,pv}",
	"",
	"{x=787,m=2655,a=1222,s=2876}",
	"{x=1679,m=44,a=2067,s=496}",
	"{x=2036,m=264,a=79,s=2244}",
	"{x=2461,m=1339,a=466,s=291}",
	"{x=2127,m=1623,a=2188,s=1013}",
}

func TestParseWorkflow(t *testing.T) {
	wf := parseWorkflow("px{a<2006:qkq,m>2090:A,rfg}")
	assert.Equal(t, Workflow{
		Name: "px",
		Stages: []Stage{
			{
				Condition:   Condition{Field: "a", Operator: '<', Value: 2006},
				Destination: "qkq",
				CatchAll:    false,
			},
			{
				Condition:   Condition{Field: "m", Operator: '>', Value: 2090},
				Destination: "A",
			},
			{
				CatchAll:    true,
				Destination: "rfg",
			},
		},
	}, wf)
}

func TestParsePart(t *testing.T) {
	part := parsePart("{x=787,m=2655,a=1222,s=2876}")
	assert.Equal(t, Part{X: 787, M: 2655, A: 1222, S: 2876}, part)
}

func TestApplyWorkflow(t *testing.T) {
	wf := parseWorkflow("in{s<1351:px,qqz}")
	part := parsePart("{x=787,m=2655,a=1222,s=2876}")
	assert.Equal(t, "qqz", part.Apply(wf))
}

func TestApplyAllWorkflows(t *testing.T) {
	wfs := parseWorkflows(example[:11])
	part := parsePart("{x=787,m=2655,a=1222,s=2876}")
	assert.Equal(t, []string{"qqz", "qs", "lnx", "A"}, part.ApplyAll(wfs))

	part = parsePart("{x=1679,m=44,a=2067,s=496}")
	assert.Equal(t, []string{"px", "rfg", "gd", "R"}, part.ApplyAll(wfs))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 19114, solvePartA(example))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 382440, solvePartA(input))
}
