package main
import "fmt"

func main()  {
	const(
		ETH = 9 -iota
		WAN
		ICX
	)
	rates := [...]float64{
		ETH: 25.5, // Etherium
		WAN: 120.5, // Wanchain
		ICX:  20, // ICX
	}
	fmt.Printf("1 BTC is  %g ETH\n", ETH)
	fmt.Printf("1 BTC is %g WAN\n",  WAN)
	fmt.Printf("1 BTC is %g ICX\n",  ICX)
	fmt.Printf("%#v\n",  rates)
}