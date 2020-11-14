package service

import (
	"context"
	"github/Ko4s/calculator_service/calc"
)

// type GreetService struct {
// 	greet.UnimplementedGreetServer
// }

// func NewGreetService() *GreetService {
// 	return &GreetService{}
// }

// func (s *GreetService) SayHello(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {

// 	name := req.GetName()

// 	res := &greet.GreetResponse{
// 		Greeting: "Hello " + name,
// 	}

// 	return res, nil
// }

type CalcService struct {
	calc.UnimplementedCalcServer
}

//CalcService constructor
func NewCalcService() *CalcService {
	return &CalcService{}
}

func (cs *CalcService) Sum(ctx context.Context, req *calc.SumRequest) (*calc.SumResponse, error) {
	//1. pobrac liczby z req
	//2. zsumować liczby
	//3. zwrocic calc.Response

	num1 := req.GetNumber1()
	num2 := req.GetNumber2()

	result := num1 + num2

	res := calc.SumResponse{
		Result: result,
	}

	return &res, nil
}
