// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type FarmInput struct {
	ID        string `json:"id"`
	Size      string `json:"size"`
	Soil      string `json:"soil"`
	ImageHash string `json:"imageHash"`
	Season    string `json:"season"`
	Owner     string `json:"owner"`
}

type HarvestInput struct {
	SeasonNumber int    `json:"seasonNumber"`
	TotalSupply  string `json:"totalSupply"`
	Price        string `json:"price"`
}

type PlantingInput struct {
	SeasonNumber  int    `json:"seasonNumber"`
	SeedUsed      string `json:"seedUsed"`
	ExpectedYield string `json:"expectedYield"`
	SeedSupplier  string `json:"seedSupplier"`
}

type PreparationInput struct {
	SeasonNumber int    `json:"seasonNumber"`
	Crop         string `json:"crop"`
	Fertilizer   string `json:"fertilizer"`
}

type SeasonInput struct {
	Token         int     `json:"token"`
	SeasonNumber  int     `json:"seasonNumber"`
	Crop          *string `json:"crop"`
	Fertilizer    *string `json:"fertilizer"`
	Seed          *string `json:"seed"`
	ExpectedYield *string `json:"expectedYield"`
	SeedSupplier  *string `json:"seedSupplier"`
	HarvestYield  *string `json:"harvestYield"`
	HarvestPrice  *string `json:"harvestPrice"`
}

type SeasonUpdateInput struct {
	Token  int    `json:"token"`
	Season string `json:"season"`
}
