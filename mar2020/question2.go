import("bufio")
import("strconv")

func contestResponse() {
  scanner := bufio.NewScanner(os.Stdin)
  nbLines := -1
  bestCount := 0
  currentCount := 0
  currentCard := 0
  for scanner.Scan() {
    line := scanner.Text()
    if nbLines == -1 {
      nbLines, _ = strconv.Atoi(line)
    } else {
      card, _ := strconv.Atoi(line)
      if card == currentCard {
          currentCount++
      } else {
          if currentCount > bestCount {
              bestCount = currentCount
          }
          currentCard = card
          currentCount = 1
      }
    }
    os.Stderr.WriteString(line)
    os.Stderr.WriteString("\n")
  }
  if currentCount > bestCount {
      bestCount = currentCount
  }
  os.Stderr.WriteString("##########################\n\n")
  fmt.Println(strconv.Itoa(bestCount))
}