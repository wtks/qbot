package conversation

import (
	"fmt"
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/conversation/respond"
	"github.com/wtks/qbot/pkg/input"
	"time"
)

var defaultErrorStep = &StepImpl{
	id: conversation.StepError,
	f: func(ctx conversation.Context, input conversation.Input) *conversation.Output {
		err := input.E()
		return respond.Message(err.Error())
	},
	end: true,
}

func BuildSpec(spec conversation.Spec) (*SpecImpl, error) {
	if len(spec.Name) == 0 {
		return nil, fmt.Errorf("name is required")
	}
	if spec.Trigger == nil {
		return nil, fmt.Errorf("trigger func is required")
	}
	if len(spec.Steps) == 0 {
		return nil, fmt.Errorf("empty step")
	}

	result := &SpecImpl{
		Name:         spec.Name,
		Steps:        map[int]*StepImpl{},
		AllowDM:      spec.AllowDM,
		AllowChannel: spec.AllowChannel,
	}
	result.Trigger = func(input input.Message) bool {
		if (result.AllowDM && input.IsDM()) || (result.AllowChannel && !input.IsDM()) {
			return spec.Trigger(input)
		}
		return false
	}

	stepMap := map[int]conversation.Step{}
	for _, s := range spec.Steps {
		if _, exists := stepMap[s.Name]; exists {
			return nil, fmt.Errorf("duplicate step: %d", s.Name)
		}
		stepMap[s.Name] = s
		stepImpl := &StepImpl{
			id:            s.Name,
			f:             s.Func,
			stampMatchers: []*stampMatcher{},
			textMatchers:  []*textMatcher{},
		}
		result.Steps[s.Name] = stepImpl
		if s.Start {
			if result.StartStep != nil {
				return nil, fmt.Errorf("multiple start step: %d, %d", result.StartStep.id, s.Name)
			}
			result.StartStep = stepImpl
		}
		if s.Name == conversation.StepError {
			result.ErrorStep = stepImpl
		}
	}
	if result.StartStep == nil {
		return nil, fmt.Errorf("start step is undefined")
	}

	for _, s := range stepMap {
		stepImpl := result.Steps[s.Name]
		if len(s.Next) == 0 {
			stepImpl.end = true
			continue
		}

		for _, flow := range s.Next {
			if _, ok := stepMap[flow.To]; !ok {
				return nil, fmt.Errorf("step %d was not defined", flow.To)
			}

			switch m := flow.Match.(type) {
			case conversation.TextMatcher:
				stepImpl.textMatchers = append(stepImpl.textMatchers, &textMatcher{
					Match: m,
					To:    result.Steps[flow.To],
				})
			case conversation.StampMatcher:
				stepImpl.stampMatchers = append(stepImpl.stampMatchers, &stampMatcher{
					Match: m,
					To:    result.Steps[flow.To],
				})
			case conversation.TimeoutMatcher:
				if stepImpl.timeoutFlow != nil {
					return nil, fmt.Errorf("step %d has multiple timeout flow", s.Name)
				}
				stepImpl.timeoutFlow = &timeoutFlow{
					Duration: time.Duration(m),
					To:       result.Steps[flow.To],
				}
			case conversation.ResultMatcher:
				stepImpl.resultMatchers = append(stepImpl.resultMatchers, &resultMatcher{
					Match: m,
					To:    result.Steps[flow.To],
				})
			default:
				return nil, fmt.Errorf("unknown flow: %v", flow)
			}
		}
	}

	if result.ErrorStep == nil {
		result.ErrorStep = defaultErrorStep
		result.Steps[conversation.StepError] = defaultErrorStep
	}

	return result, nil
}
