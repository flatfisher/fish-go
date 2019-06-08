package main

func main(){
	car := NewBuilder().Color(BlueColor).Wheels(SteelWheels).TopSpeed(KPH).Build()
	car.Drive()
	car.Stop()
}