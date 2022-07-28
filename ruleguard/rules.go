package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func removeBadStringAppend(m dsl.Matcher) {
	// A chain that starts with Match() and ends with Report() call called a *rule*.
	// Therefore, a minimal *rule* consists of Match+Report call.
	//
	// Match() matches the AST using the gogrep pattern string;
	// Report() prints the specified message when this rule is matched.
	m.Import("fmt")

	m.Match(
		`$x + $v`,
		`$x += $v`,
	).Where(
		m["x"].Type.Is(`string`),
	).Report(`Appending strings with plus operator found!`)
}

func enforceFactoryPattern(m dsl.Matcher) {

	m.Match(
		`MyObject{$*_}`,
		`new(MyObject)`,
	).Report("Only use factory method 'MakeMyObject' to initialize object.")
}


func detectRecursive(m dsl.Matcher) {

	// detect if function has return
	hasReturn := func(v dsl.Var) bool {
		// query selected text with regexp statement
		return v.Text.Matches(`if([\W,\w]+){([\W,\w]+)return([\W,\w]+)}`)
	}

	m.Match(
		`func $name($*vars){
			$*body
			$name($v)
			$*trailing
		}`,
	).Where(
		!hasReturn(m["body"]),
	).Report("Function $name will run until the end of time.")
}