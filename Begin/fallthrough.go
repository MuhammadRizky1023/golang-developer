

func main() {
 i:= 142;
switch {
	case i> 100:
		fmt.Println("Big");
		fallthrough
	 case i > 0:
		fmt.Println("positive");
		fallthrough
  default:
	fmt.Println("number")
  }
}