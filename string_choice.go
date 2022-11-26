package eflags

import (
	//"errors"
	"fmt"
	"strings"
)

type stringChoiceValue struct {
	options	[]string
	value	*string
	ci		bool
}


//////////////////////////////////////////////////////////////////////////////
// implement the Value interface

func (scv *stringChoiceValue) Type() string {
	return "stringChoice"
}

func (scv *stringChoiceValue) Value() string {
	return *scv.value
}

func (scv *stringChoiceValue) String() string {
	return *scv.value
}

func (scv *stringChoiceValue) Set(val string) error {

	if nil == scv.value {
		return fmt.Errorf("StringChoiceValue::Set(): nil value")
	}

	for _,c := range scv.options {
		if scv.ci && strings.EqualFold(val,c) {
			*scv.value = c		// pre-defined is "canonical"
			return nil
		} else if val == c {
			*scv.value = c
			return nil
		}
	}	
	return fmt.Errorf(`Value "%s" not within allowed choices`, val)
}


//////////////////////////////////////////////////////////////////////////////
func StringChoice(valp *string, defval string, choices []string) *stringChoiceValue {
	if nil == valp {
		var v string
		return &stringChoiceValue{value: &v}
	}
	scv := &stringChoiceValue{value: valp, options:choices}
	*scv.value=defval	// !
	return scv
}

func StringChoiceCase(valp *string, defval string, choices []string) *stringChoiceValue {
	if nil == valp {
		var v string
		return &stringChoiceValue{value: &v, ci: true}
	}
	scv := &stringChoiceValue{value: valp, options:choices, ci:true}
	*scv.value=defval	// !
	return scv
}
