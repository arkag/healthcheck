package healthcheck_fetch

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloWorld(t *testing.T) {
    text := "hello world"
    want := regexp.MustCompile(`\b`+text+`\b`)
    msg, err := healthcheck_fetch()
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`healthcheck_fetch = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}