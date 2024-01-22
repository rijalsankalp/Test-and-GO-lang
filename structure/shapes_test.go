package structure

import "testing"

func TestShapes(t *testing.T) {

	testarea := []struct{
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 2.0}, want:20.0},
		{name: "Circle", shape: Circle{10.0}, want:314.1592653589793},
		{name: "Triangle", shape: Triangle{4.0,3.0}, want: 6.0},
	}

	for _,ta := range testarea{
		t.Run(ta.name, func(t *testing.T) {
			hasArea := ta.shape.Area()
			if hasArea != ta.want{
				t.Errorf("%#v want %g, and got %g",ta.shape, ta.want, hasArea)
			}
		})
	}


	testperimeter := []struct{
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 2.0}, want:24.0},
		{name: "Circle", shape: Circle{3.0}, want:18.84955592153876},
		{name: "Triangle", shape: Triangle{4.0,3.0}, want: 6.0},
	}

	for _,tp := range testperimeter{
		t.Run(tp.name, func(t *testing.T) {
			hasPeri := tp.shape.Perimeter()
			if hasPeri != tp.want{
				t.Errorf("%#v want %g, and got %g", tp.shape, tp.want, hasPeri)
			}
		})
	}

}