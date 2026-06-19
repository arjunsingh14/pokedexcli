package main

import "testing"

  func TestCleanInput(t *testing.T) {
        cases := []struct {
                input  string
                output []string
        }{
                {
                        input:  "hello       world      ",
                        output: []string{"hello", "world"},
                },
        }

        for _, c := range cases {
                actual := cleanInput(c.input)
                if len(actual) != len(c.output) {
                        t.Errorf("lengths don't match: got %d, want %d", len(actual), len(c.output))
                        continue
                }
                for i := range actual {
                        word := actual[i]
                        expectedWord := c.output[i]
                        if word != expectedWord {
                                t.Errorf("Word and expected word do not match: got %q, want %q", word, expectedWord)
                        }
                }
        }
  }