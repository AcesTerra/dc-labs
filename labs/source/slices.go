package main

import "golang.org/x/tour/pic"
//import "fmt"

func Pic(dx, dy int) [][]uint8 {
		pixel := make([][]uint8, dy)
		 	data := make([]uint8, dx)
				//fmt.Println(len(data))
					for i := range pixel {
						  		for j := range data {
									   			//data[j] = uint8((i+j))
															if i+j > 128 && j > 128{
																				data[j] = uint8((255))
																							}
																							  		}
																									  		pixel[i] = data
																											 	}
																												 	return pixel
																												}

																												func main() {
																														pic.Show(Pic)
																													}
