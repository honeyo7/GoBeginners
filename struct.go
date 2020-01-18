package main

import "fmt"

type car struct  {
	gas_pedal uint16 // min 0 to max 65535
	brake_pedal uint16
	steering_wheel int16// -32K to 32K
	top_speed_kmh float64
}

type supercar struct { //embedded struct
	car
	can_fly bool
}

func main() {

	a_car := car{gas_pedal:22341,
				brake_pedal:0,
				steering_wheel: 12561,
				top_speed_kmh:225.0,
			}
				
	fmt.Println(a_car.gas_pedal)


	s_car := supercar{
		car: car{
			gas_pedal: 22341,
			brake_pedal: 0,
			steering_wheel: 21561,
			top_speed_kmh: 1225.0,

		},
		can_fly: true,

	}

	fmt.Println(s_car.top_speed_kmh)

}