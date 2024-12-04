package main

import(
	f "fmt"
)

func main(){
	provincias := []string {
		"Buenos Aires", "Catamarca", "Chaco", "Chubut", "Córdoba", "Corrientes", "Entre Ríos", "Formosa", "Jujuy", "La Pampa", "La Rioja", "Mendoza", "Misiones", "Neuquén", "Río Negro", "Salta", "San Juan", "San Luis", "Santa Cruz", "Santa Fe", "Santiago del Estero", "Tierra del Fuego, Antártida e Islas del Atlántico Sur", "Tucumán"}

		f.Println("La longitud es :" , len(provincias))
		f.Println("La capacidad es :" , cap(provincias))

		f.Println("--------------------")
		f.Println("Usando Range")
		for indice , valor :=  range provincias{
			f.Printf("Indice %d ,  valor %s \n" , indice , valor )
		}
		
		f.Println("--------------------")
		f.Println("Sin usar el Range")

		for i := 0 ; i < len(provincias) ; i++{
			f.Println("Provincia: " , provincias[i] , "en indice : " , i)
		}

}