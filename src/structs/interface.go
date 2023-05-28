type PrintInfo interface {
	getMessage() string
	setMessage()
}
//this way can implement printInfo interface
func (p PrintInfo)  {
	fmt.Println(p.getMessage())
}