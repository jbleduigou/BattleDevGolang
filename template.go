import("bufio")
import("strconv")
func contestResponse() {
	scanner := bufio.NewScanner(os.Stdin)
	nbLines := -1
	result := ""
	for scanner.Scan() {
		line := scanner.Text()
		if nbLines == -1 {
		  nbLines, _ = strconv.Atoi(line)
		} else {
		    result = line
		}
		os.Stderr.WriteString(line)
        os.Stderr.WriteString("\n")
	}
	os.Stderr.WriteString("##########################\n\n")
	fmt.Println(result)
}
