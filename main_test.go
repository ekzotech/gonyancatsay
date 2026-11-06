package main

import (
	"strings"
	"testing"
)

func TestSpeechBubble(t *testing.T) {
	got := speechBubble("Привет")
	want := " ________\n< Привет >\n --------"
	if got != want {
		t.Fatalf("speechBubble() =\n%s\nwant:\n%s", got, want)
	}
}

func TestSpeechBubbleMultiline(t *testing.T) {
	got := speechBubble("Nyan\nкот")
	want := " ______\n/ Nyan \\\n\\ кот  /\n ------"
	if got != want {
		t.Fatalf("speechBubble() =\n%s\nwant:\n%s", got, want)
	}
}

func TestRenderRainbow(t *testing.T) {
	got := render("Nyan", true)
	if !strings.HasPrefix(got, "\x1b[31m ______\x1b[0m\n") {
		t.Fatalf("rainbow output does not start with a colored border: %q", got)
	}
	if !strings.Contains(got, "\x1b[34m         \\\x1b[0m") {
		t.Fatalf("rainbow output does not contain colored mascot: %q", got)
	}
	if strings.Count(got, "\x1b[0m") != len(strings.Split(got, "\n")) {
		t.Fatal("not every line resets its color")
	}
}
