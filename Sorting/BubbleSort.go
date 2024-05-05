package Sorting

func Burbujeo(arreglo []int) []int {
	for i:=0; i < len(arreglo); i+=1{
    for j := 0; j < len(arreglo)-1-i; j+=1{
      if arreglo[j] > arreglo[j+1]{
        tmp := arreglo[j]
        arreglo[j] = arreglo[j+1]
        arreglo[j+1] = tmp
      }
    }

	}

  return arreglo
}
