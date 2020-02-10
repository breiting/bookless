package main

func main() {

	conf := GetDefaultConfig()

	a := &App{}
	a.Initialize(conf)
	a.Run(":8080")
}
