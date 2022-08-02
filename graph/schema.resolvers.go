package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"errors"
	
	

	"github.com/Roshank/Demo/graph/generated"
	"github.com/Roshank/Demo/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Addition is the resolver for the addition field.
func (r *queryResolver) Addition(ctx context.Context, input *model.AdditionRequest) (*model.AdditionResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.AdditionResponse
	num3 = num1 + num2
	mod = model.AdditionResponse{Number3: (&num3)}
	return &mod, nil
}

// Subtraction is the resolver for the subtraction field.
func (r *queryResolver) Subtraction(ctx context.Context, input *model.SubtractionRequest) (*model.SubtractionResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.SubtractionResponse
	num3 = num1 - num2
	mod = model.SubtractionResponse{Number3: (&num3)}
	return &mod, nil
}

// Multiplication is the resolver for the multiplication field.
func (r *queryResolver) Multiplication(ctx context.Context, input *model.MultiplicationRequest) (*model.MultiplicationResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.MultiplicationResponse
	num3 = num1 * num2
	mod = model.MultiplicationResponse{Number3: (&num3)}
	return &mod, nil
}

// Division is the resolver for the division field.
func (r *queryResolver) Division(ctx context.Context, input *model.DivisionRequest) (*model.DivisionResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.DivisionResponse
	num3 = num1 / num2
	mod = model.DivisionResponse{Number3: (&num3)}
	return &mod, nil
}

// Calculation is the resolver for the calculation field.
func (r *queryResolver) Calculation(ctx context.Context, input *model.CalculationRequest) (*model.CalculationResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	// var operator string
	var num1 int
	var num2 int
	var num3 int
	operator:= *input.Operator
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.CalculationResponse
	if operator == "+"{
		num3 = num1 + num2
	}
	if operator == "-"{
		num3 = num1 - num2
	}
	if operator == "*"{
		num3 = num1 * num2
	}
	if operator == "/"{
		// if err!= nil{
		// 	fmt.Println("cant divide")
		// }
		if(num2 == 0){
			return nil,errors.New("can't divide by zero")
		}
		num3 = num1 / num2
	}
	mod = model.CalculationResponse{Number3: (&num3)}
	return &mod, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
