package main

import "fmt"

type car struct  {
	car_name string
	gas_pedal uint16 // min 0 to max 65535
	brake_pedal uint16
	steering_wheel int16// -32K to 32K
	top_speed_kmh float64
}

type supercar struct { //embedded struct
	car
	can_fly bool
}

func (c car) top_speed(){
	fmt.Println(c.car_name, "top speed is", c.top_speed_kmh)
}

func (sc supercar) top_speed(){
	fmt.Println(sc.car_name, "top speed is", sc.top_speed_kmh)
}

type four_wheeler interface{ // Interface say, hey baby if you got these methods then you are also my type
	top_speed()
}

func foo (f four_wheeler){  
	f.top_speed()
}

func main() {

	a_car := car{ car_name: "abc car",
		gas_pedal:22341,
				brake_pedal:0,
				steering_wheel: 12561,
				top_speed_kmh:225.0,
			}
				
	fmt.Println(a_car.gas_pedal)


	s_car := supercar{
		car: car{car_name: "xyz car",
			gas_pedal: 22341,
			brake_pedal: 0,
			steering_wheel: 21561,
			top_speed_kmh: 1225.0,

		},
		can_fly: true,

	}

	fmt.Println(s_car.top_speed_kmh)

	foo(a_car)
	foo(s_car)

}