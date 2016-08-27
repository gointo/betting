### `team` from bet splitted on ` ` can be replace by the team matching in team check

## Bet checks

```
- `\n`(`team`|both teams) to score * [in (1st|2nd|first|second) half]
- `\n`[`team`] (over|under) `float` * goals  
- '\n'`team` no clean sheet
- '\n'`team` win(s)
- '\n'[`team`] (fulltime|halftime| 30 minute) result [`team`]
- '\n'`team` to `int` corner(s)
- '\n'`team` [to score their] *((1st|2nd|3rd|4th|5th|6th|7th|8th|9th|first|second|third|fourth|fifth) *goal *(before|after|between) *int('|:int)
- '\n'`team` or draw
```

## Team checks
```
- `team` vs `team`
- `team` game
```

```go

enum Bets {
  TOSCORE iota
  OVERUNDER
  CLEANSHEET
  WINNER
  RESULT
  CORNERS
  GOALBEFORE
  DOUBLECHANCE
}

enum When {
  BEFORE iota,
  BETWEEN
  AFTER
}

// First team is the one to bet on if Specific is true
type Bet struct {
  Team []string,
  Specific bool,
  Type string,
  Stake float,
}

type ToScore struct {
  Bet,
  Half uint8,
}

type OverUnder struct {
  Bet,
  Goals float,
}

type CleanSheet struct {
  Bet,
  Is bool,
}

type Winner struct {
  Bet,
}

type Result struct {
  Bet,
  Time uint8,
}

type Corners struct {
  Bet,
  Number uint8
}

Type GoalPeriod struct }
  Bet,
  Goal uint8,
  When int,
  Period int,
}

impl Bets string()
impl Bets init()
```
