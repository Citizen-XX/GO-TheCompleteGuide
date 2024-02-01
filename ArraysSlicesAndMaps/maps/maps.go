package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Amazon": "https://www.aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon"])

	websites["LinkedIn"] = "http://www.linkedin.com"
	fmt.Println(websites)
}