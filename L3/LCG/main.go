package main

import (
	"fmt"
	"math"
)

func egcd(a,b int) (int,int,int){
	if a == 0{
		return b,0,1
	}else{
		g,y,x :=egcd(b%a,a)
		return g,x-(b/a)*y,y
	}
}

func mod_inverse(a,m int) (int,error){
	g,x,_ := egcd(a,m)
	if g!=1{
		return 0,fmt.Errorf("mod inverse does not exist for 1")
	}else{
		return x%m,nil
	}
}

func find_a(x[3]int,m int) (int,error){
	inverse_m,err := mod_inverse(x[1]-x[0],m)
	if err!=nil{
		return 0,err
	}
	a := (x[2] - x[1]) * inverse_m%m
	return a,nil
}

func find_c(x [3]int,a,m int) int{
	c:= (x[1] - x[0] * a) % m
	return c
}

func lcg(xPrev,a,c,m int) int{
	return (xPrev*a+c)%m
}

func main(){
	x :=[3]int{-428720726,2053624065,974386796}
	m:=int(math.Pow(2,32))
	a,err := find_a(x,m)
	if err!=nil{
		print(err.Error())
	}else{
		fmt.Printf("a:%v\n",a )
	}
	c:= find_c(x,a,m)
	fmt.Printf("c:%v\n",c )
	xNext := lcg(x[2],a,c,m)
	fmt.Printf("Next x :%v\n",xNext )

	/*
	//x :=[]int{-1517823580,101734579,-1206532490,-1541207203,-1529703144,299898263}
	m:=int(math.Pow(2,32))
	c := 1013904223
	a := 1664525
	f := lcg(1818388024,a,c,m)

	print(f-m)
	/*
	for a:=0;a<m;a++{
			x1 := (x[0]*a +c)%m
			//print(x1)
			if x1!=x[1]{
				continue
			}
			x2 := (x[3]*a +c)%m
			x3 := (x[4]*a +c)%m
			x4 := (x[5]*a +c)%m
			if  x2 ==x[2] && x3==x[3] && x4==x[4]{
				fmt.Printf("m=%d, a=%d, c=%d\n", m, a, c)
		}
	}
	 */

}
