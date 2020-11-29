import("bufio")
import("strconv")
import("strings")
// Plus d'infos sur https://github.com/jbleduigou/BattleDevGolang
func contestResponse() {
  scanner := bufio.NewScanner(os.Stdin)
  nbLines := -1
  result := "OK"
  nightly := 0.0
  total := 0.0
  for scanner.Scan() {
    line := scanner.Text()
    if nbLines == -1 {
      nbLines, _ = strconv.Atoi(line)
    } else {
        s := strings.Split(line, ":")
        hour, _ := strconv.Atoi(s[0])
        if hour < 8 {
            nightly = nightly + 1.0
        }
        if hour >= 20 {
            nightly = nightly + 1.0
        }
        total = total + 1.0
    }
    os.Stderr.WriteString(line)
    os.Stderr.WriteString("\n")
  }
  if nightly / total > .5 {
      result = "SUSPICIOUS"
  }
  os.Stderr.WriteString("##########################\n\n")
  fmt.Println(result)
}