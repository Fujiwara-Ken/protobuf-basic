package main

import (
	"protobuf-lesson/pb"
	"github.com/golang/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	// シリアライズ

	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Cant serialize", err)
	// }

	// if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
	// 	log.Fatalln("Cant write", err)
	// }

	// in, err := ioutil.ReadFile("test.bin")
	// if err != nil {
	// 	log.Fatalln("Cant read file", err)
	// }

	// readEmployee := &pb.Employee{}

	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("cant deserialize", err)
	// }

	// fmt.Println(readEmployee)

	// json変換
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if != err nil {
		log.Fatalln("cant marshal to json", err)
	}

	fmt.Pringln(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("cant unmarshal from json", err)
	}

	fmt.Println(readEmployee)
}
