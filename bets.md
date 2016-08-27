### `team` from bet splitted on ` ` can be replace by the team matching in team check

## Bet checks

```
- (`team`|both teams) *to *score * [in (1st|2nd|first|second) half]
- [`team`] *(over|under) *`float` * goals  
- `team` *no *clean sheet
- `team` *win(s)
- [`team`] *(fulltime|halftime| 30 minute) *result *[`team`]
- `team` *to *`int` *corner(s)
```

## Team checks
```
- `team` vs `team`
- `team` game
```

```go

enum {
  TOSCORE iota
  OVERUNDER
  CLEANSHEET
  RESULT
  CORNERS
}

// First team is the one to bet on if Specific is true
Type Bets truct {
  Team []string,
  Specific bool,
  Type string,
  Stake float,
}

impl Bets string()
impl Bets init()
```
