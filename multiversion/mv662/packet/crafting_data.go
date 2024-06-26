package packet

import (
	"fmt"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// CraftingData is sent by the server to let the client know all crafting data that the server maintains. This
// includes shapeless crafting, crafting table recipes, furnace recipes etc. Each crafting station's recipes
// are included in it.
type CraftingData struct {
	// Recipes is a list of all recipes available on the server. It includes among others shapeless, shaped
	// and furnace recipes. The client will only be able to craft these recipes.
	Recipes []protocol.Recipe
	// PotionRecipes is a list of all potion mixing recipes which may be used in the brewing stand.
	PotionRecipes []protocol.PotionRecipe
	// PotionContainerChangeRecipes is a list of all recipes to convert a potion from one type to another,
	// such as from a drinkable potion to a splash potion, or from a splash potion to a lingering potion.
	PotionContainerChangeRecipes []protocol.PotionContainerChangeRecipe
	// MaterialReducers is a list of all material reducers which is used in education edition chemistry.
	MaterialReducers []protocol.MaterialReducer
	// ClearRecipes indicates if all recipes currently active on the client should be cleaned. Doing this
	// means that the client will have no recipes active by itself: Any CraftingData packets previously sent
	// will also be discarded, and only the recipes in this CraftingData packet will be used.
	ClearRecipes bool
}

// ID ...
func (*CraftingData) ID() uint32 {
	return packet.IDCraftingData
}

func (pk *CraftingData) Marshal(io protocol.IO) {
	w := io.(*protocol.Writer)
	protocol.FuncSlice(io, &pk.Recipes, func(x *protocol.Recipe) {
		var recipeType int32
		if !lookupRecipeType(*x, &recipeType) {
			w.UnknownEnumOption(fmt.Sprintf("%T", *x), "crafting recipe type")
		}
		w.Varint32(&recipeType)
		(*x).Marshal(w)
	})
	protocol.Slice(io, &pk.PotionRecipes)
	protocol.Slice(io, &pk.PotionContainerChangeRecipes)
	protocol.FuncSlice(io, &pk.MaterialReducers, io.MaterialReducer)
	io.Bool(&pk.ClearRecipes)
}

// lookupRecipeType looks up the recipe type for a Recipe. False is returned if
// none was found.
func lookupRecipeType(x protocol.Recipe, recipeType *int32) bool {
	switch x.(type) {
	case *protocol.ShapelessRecipe:
		*recipeType = protocol.RecipeShapeless
	case *protocol.ShapedRecipe, *ShapedRecipe:
		*recipeType = protocol.RecipeShaped
	case *protocol.FurnaceRecipe:
		*recipeType = protocol.RecipeFurnace
	case *protocol.FurnaceDataRecipe:
		*recipeType = protocol.RecipeFurnaceData
	case *protocol.MultiRecipe:
		*recipeType = protocol.RecipeMulti
	case *protocol.ShulkerBoxRecipe:
		*recipeType = protocol.RecipeShulkerBox
	case *protocol.ShapelessChemistryRecipe:
		*recipeType = protocol.RecipeShapelessChemistry
	case *protocol.ShapedChemistryRecipe, *ShapedChemistryRecipe:
		*recipeType = protocol.RecipeShapedChemistry
	case *protocol.SmithingTransformRecipe:
		*recipeType = protocol.RecipeSmithingTransform
	case *protocol.SmithingTrimRecipe:
		*recipeType = protocol.RecipeSmithingTrim
	default:
		return false
	}
	return true
}
