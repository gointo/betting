package main

import (
  "regexp"
  "log"
  "strings"
)

func checkBet(res []string, num int) bool {
  i := 0
  for i < len(res) {
    if len(res) == 0 {
      return false
    }
    i++
  }
  if i != num {
    return false
  }
  return true
}

func checkToScore(body string) ([]string, string, string) {
  re := regexp.MustCompile("\n(both teams|.*) to score(.*)(@ .*)\n")
  res := re.FindStringSubmatch(body)
  log.Printf("bet is score: %q", res)
  if !checkBet(res, 4) {
    log.Printf("This bet is not a ToScore")
    return nil, "", ""
  }
  log.Printf("This bet is a ToScore: %q", res)
  return []string{res[1]}, "TO SCORE", res[2]
}

func getGame(game []string, msg string) ([]string, bool) {
  re := regexp.MustCompile("\n(.*) vs ( .*)\n")
  res := re.FindStringSubmatch(msg)
  if checkBet(res, 3) {
    log.Printf("Both team specified")
    if strings.Contains(game[0], res[1]) {
      return res[1:], true
    } else if strings.Contains(game[0], res[2]) {
      return []string{res[2], res[1]}, true
    }
  }
  re = regexp.MustCompile("\n(.*) game\n")
  res2 := re.FindStringSubmatch(msg)
  if checkBet(res2, 2) && res[0] == "" {
    if strings.Contains(game[0], res2[1]) {
      return []string{res2[1]}, true
    } else if strings.Contains(res2[1], game[0]) {
      return game, true
    }
  if res2[1] != "" {
    return []string{res2[1]}, true
  } else if res[1] != "" {
    return res[1:], false
  } else {
    return game, true
  }
  }
  return game, true
}

func fmtBet(game []string, trgt bool) string {
  if trgt {
	  return game[0] + " vs " + game[1] + "\nSPECIFIC: true"
  }
  return game[0] + " GAME \nSPECIFIC: false"
}

func getBet(body string) (string, string, string) {
  game, bet, obj := checkToScore(body)
  if len(game) > 1 && game[0] != "" && bet != "" {
    game, specific := getGame(game, body)
    return fmtBet(game, specific), bet, obj
  }
  return "", "", ""
}

func getData(created string, msg string) string {
  game, bet, obj := getBet(msg)
  // TODO getStake(body string) string
  if bet != "" {
    return bet + "\n" + game + "\nDETAILS: " + obj + "\n" + "DATE: " + created + "\n"
  }
  return ""
}
