package service

import (
	"errors"
	"main/pkg/request"
	"main/pkg/response"

	"github.com/gin-gonic/gin"
)

func Service(c *gin.Context, reqBody request.Request) (response.Response, error) {
	respChan := make(chan response.Response)
	errChan := make(chan error)
	go func() {
		resp, err := FindPair(reqBody)
		if err != nil {
			errChan <- err
		} else {
			respChan <- resp
		}

	}()

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return response.Response{}, err
	}

}

func FindPair(reqBody request.Request) (response.Response, error) {
	combinationCount := make(map[int]int)
	numCount := []response.Solution{}
	for i, v := range reqBody.Numbers {
		rem := reqBody.Target - v
		if idx, exists := combinationCount[rem]; exists {
			numCount = append(numCount, response.Solution{idx, i})
		}

		combinationCount[v] = i
	}

	if len(numCount) < 1 {
		return response.Response{}, errors.New("target is not matching with the sum of any two elements in the array")
	}

	return response.Response{
		Solution: numCount,
	}, nil
}
