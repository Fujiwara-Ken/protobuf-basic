package main

import "protobuf-lesson/pb"

func main(){
	employee := &pb.Employee {
		Id: 1,
		Name: "Suzuki",
		Email: "test@example.com",
		Occupation: pb.Occupation_ENFINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project: map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: & pb.Employee_Text {
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year: 2000,
			Month: 1,
			Day: 1,
		},
	} 
}