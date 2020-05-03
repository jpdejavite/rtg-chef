package models

type RecipeAppQueries struct {
	List []*RecipeList `json:"list"`
}

type RecipeList struct {
	Total   int       `json:"total"`
	Recipes []*Recipe `json:"recipes"`
}

type Recipe struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
