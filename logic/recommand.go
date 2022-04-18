package logic

import (
	"resturant_recommand_proxy/client"
	"strconv"
)

type orderRate struct {
	productName string
	rate        float64
}

func getProductRank(username string) ([]orderRate, error) {
	res := []orderRate{}
	val, err := client.Rdb.Get(username).Result()
	if err != nil {
		return res, err
	}
	count, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return res, err
	}

	products := []string{}
	for _, p := range products {
		val1, err := client.Rdb.Get(p).Result()
		if err != nil {
			return res, err
		}
		count1, err := strconv.ParseFloat(val1, 64)
		if err != nil {
			return res, err
		}
		pRate := orderRate{
			productName: p,
			rate:        0.5*count + 0.5*count1,
		}
		res = append(res, pRate)
	}
	quickSort(res, 0, len(res)-1)
	return res, nil
}

func quickSort(rates []orderRate, left, right int) {
	if left < right {
		index := getIndex(rates, left, right)
		quickSort(rates, left, index-1)
		quickSort(rates, index+1, right)
	}
}

func getIndex(rates []orderRate, left, right int) int {
	temp := rates[left]
	for left < right {
		for left < right && rates[right].rate <= temp.rate {
			right--
		}
		if left < right {
			rates[left] = rates[right]
		}
		for left < right && rates[left].rate >= temp.rate {
			left++
		}
		if left < right {
			rates[right] = rates[left]
		}
	}
	rates[left] = temp
	return left
}
