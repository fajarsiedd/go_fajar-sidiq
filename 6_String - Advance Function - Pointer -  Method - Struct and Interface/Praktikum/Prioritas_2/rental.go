package main

import (
	"rental/customer"
	"rental/vehicle"
)

func main() {
	trx := customer.Transaction{}

	trx.AddCustomer("Fajar", "Jl. Cibuntu")
	trx.AddCustomer("Asep", "Jl. Abc")
	trx.AddCustomer("Rosyid", "Jl. Def")

	trx.GetCustomers()

	trx.RentVehicle("Asep", vehicle.Vehicle{Brand: "Honda", Type: "SUV", RentCost: 550000})
	trx.RentVehicle("Fajar", vehicle.Vehicle{Brand: "BMW", Type: "Sedan", RentCost: 350000})
	trx.RentVehicle("Fajar", vehicle.Vehicle{Brand: "Toyota", Type: "Sedan", RentCost: 450000})
	trx.RentVehicle("Fajar", vehicle.Vehicle{Brand: "BMW", Type: "MVP", RentCost: 450000})

	trx.ReturnVehicle("Fajar", "Toyota Sedan")

	trx.GetRentedVehicles("Fajar")
}
