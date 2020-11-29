import("bufio")
import("regexp")
import("strconv")
func contestResponse() {
  scanner := bufio.NewScanner(os.Stdin)
  nbLines := -1
  result := 0
  re := regexp.MustCompile(`(.)*([0-9]{5})`)
  for scanner.Scan() {
    line := scanner.Text()
    if nbLines == -1 {
      nbLines, _ = strconv.Atoi(line)
    } else {
      if (re.MatchString(line)) {
        result++
      }
    }
    os.Stderr.WriteString(line)
    os.Stderr.WriteString("\n")
  }
  os.Stderr.WriteString("##########################\n\n")
  fmt.Println(strconv.Itoa(result))
}