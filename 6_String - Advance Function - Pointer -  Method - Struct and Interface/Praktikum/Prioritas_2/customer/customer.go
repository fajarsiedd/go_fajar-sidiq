package customer

import (
	"fmt"
	"rental/vehicle"
	"slices"
	"strings"
)

type Customer struct {
	Name           string
	Address        string
	RentedVehicles []vehicle.Vehicle
}

type Transaction struct {
	Customers []Customer
}

func (trx *Transaction) AddCustomer(name, address string) {
	newCustomer := Customer{Name: name, Address: address}
	trx.Customers = append(trx.Customers, newCustomer)

	slices.SortFunc(trx.Customers, func(a, b Customer) int {
		return strings.Compare(a.Name, b.Name)
	})
}

func (trx Transaction) GetCustomers() {
	fmt.Println()
	fmt.Println("===== DATA PELANGGAN =====")
	fmt.Println("Nama\t\tAlamat")
	fmt.Println("--------------------------")
	for _, v := range trx.Customers {
		fmt.Printf("%s\t\t%s\n", v.Name, v.Address)
	}
	fmt.Println("==========================")
	fmt.Println()
}

func (trx *Transaction) RentVehicle(cust_name string, vhc vehicle.Vehicle) {
	index, found := slices.BinarySearchFunc(trx.Customers, Customer{Name: cust_name}, func(a, b Customer) int {
		return strings.Compare(a.Name, b.Name)
	})

	if found {
		trx.Customers[index].RentedVehicles = append(trx.Customers[index].RentedVehicles, vhc)

		slices.SortFunc(trx.Customers[index].RentedVehicles, func(a, b vehicle.Vehicle) int {
			return strings.Compare(a.Brand+" "+a.Type, b.Brand+" "+b.Type)
		})

		fmt.Printf("[log] %s telah menyewa kendaraan merk %s tipe %s\n", cust_name, vhc.Brand, vhc.Type)
	} else {
		fmt.Println("[log] Pelanggan tidak ditemukan!")
	}
}

func (trx *Transaction) ReturnVehicle(cust_name string, vhc_name string) {
	c_index, c_found := slices.BinarySearchFunc(trx.Customers, Customer{Name: cust_name}, func(a, b Customer) int {
		return strings.Compare(a.Name, b.Name)
	})

	if c_found {
		temp := strings.Split(vhc_name, " ")
		target := vehicle.Vehicle{Brand: temp[0]}
		if len(temp) > 1 {
			target.Type = temp[1]
		}

		v_index, v_found := slices.BinarySearchFunc(trx.Customers[c_index].RentedVehicles, target, func(a, b vehicle.Vehicle) int {
			return strings.Compare(a.Brand+" "+a.Type, b.Brand+" "+b.Type)
		})

		if v_found {
			trx.Customers[c_index].RentedVehicles = slices.Delete(trx.Customers[c_index].RentedVehicles, v_index, v_index+1)
			fmt.Printf("[log] %s telah mengembalikan kendaraan %s\n", cust_name, vhc_name)
		} else {
			fmt.Println("[log] Kendaraan tidak ditemukan!")
		}
	} else {
		fmt.Println("[log] Pelanggan tidak ditemukan!")
	}
}

func (trx Transaction) GetRentedVehicles(cust_name string) {
	index, found := slices.BinarySearchFunc(trx.Customers, Customer{Name: cust_name}, func(a, b Customer) int {
		return strings.Compare(a.Name, b.Name)
	})

	if found {
		fmt.Println()
		fmt.Println("======== DATA SEWA =======")
		fmt.Println("Nama Pelanggan: ", cust_name)
		fmt.Println("--------------------------")
		fmt.Println("Merk\tTipe\tHarga Sewa")
		fmt.Println("--------------------------")
		for _, v := range trx.Customers[index].RentedVehicles {
			fmt.Printf("%s\t%s\t%d\n", v.Brand, v.Type, v.RentCost)
		}
		fmt.Println("==========================")
		fmt.Println()
	} else {
		fmt.Println("[log] Pelanggan tidak ditemukan!")
	}
}
