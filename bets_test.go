package main

import (
  "reflect"
  "testing"
)

var toscorebetsTests = []struct {
  body string
  game []string
  bet string
  details string
}{
  {"blabla", nil, "", ""},
  {"\n\nlyon to score the first goal @ 1.40\n\nstake 1u", []string{"lyon "}, "TO SCORE", " the first goal"},
  {"waza u42 vs st johnstone u20\n\nwaza u42 race to 9 corners @ 1.42\n\nstake only 0.5u. as they wont push enough if they score.", nil, "", ""},
  {"denver vs aldershot\n\ndenver to score in 2nd half @ 1.21\n\nstake 1u", []string{"denver "}, "TO SCORE", " in 2nd half"},
  {"ashton kutsher vs lord king\n\nlord to score in 2nd half @ 1.11\n\nstake 1u.", []string{"lord "}, "TO SCORE", " in 2nd half"},
  {"silver plate to score the 6th goal @ 3.50\n\nstake 0.5u", []string{"silver plate "}, "TO SCORE", " the 6th goal"},
  {"dominican republic vs puerto rico\n\ndominican to score in 2nd half @ 1.53\n\nstake 1u.", []string{"silver plate "}, "TO SCORE", " the 6th goal"},
}

func TestToScoreBets(t *testing.T) {
  for _, test := range toscorebetsTests {
    game, bet, details := checkToScore(test.body)
    if !reflect.DeepEqual(game, test.game) || bet != test.bet || details != test.details {
      t.Errorf("Game: '%s' | Bet: '%s' | Details: '%s'", game, bet, details)
    }
  }
}
