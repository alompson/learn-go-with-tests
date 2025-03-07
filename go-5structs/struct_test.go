package structs

import (
	"testing"
)

// func TestPerimeter(t *testing.T){
// 	rect := Rectangle{10.0, 10.0}
// 	got := Perimeter(rect)
// 	want := 40.0
// 	checkValue(t, got, want)
// }

func TestArea(t *testing.T){

	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "rectangle", shape: Rectangle{Width:12, Height:6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 5}, want: 25.0},
	}
	
	for _,tt := range areaTests{
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T){
			got := calculateArea(tt.shape)
			checkValue(t, tt.shape, got, tt.want)
		})
	}
}


func checkValue(t testing.TB, shape Shape, got, want float64){
	t.Helper()
	if got != want{
		t.Errorf("%v case: got %g but want %g", shape, got, want)
	}
}