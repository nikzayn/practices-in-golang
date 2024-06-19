package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type Product struct {
		Price       float64 // 8 bytes
		ID          int32   // 4 bytes
		Title       string  // 16 bytes
		IsAvailable bool    // 1 byte
	}

	type OptimizedProduct struct {
		IsAvailable bool    // 1 byte
		ID          int32   // 4 bytes
		Price       float64 // 8 bytes
		Title       string  // 16 bytes
	}

	// Stacks Product list Restucturing
	type Badge struct {
		Type     string `json:"type"`
		Location string `json:"location"`
		Image    string `json:"image,omitempty"`
		Text     string `json:"text,omitempty"`
	}

	type Swatch struct {
		SwatchImage                 string   `json:"swatchImage,omitempty"`
		ProductImage                []string `json:"productImage,omitempty"`
		Family                      string   `json:"family,omitempty"`
		Name                        string   `json:"name,omitempty"`
		GenericLabel                string   `json:"genericLabel,omitempty"`
		ChoiceLabel                 string   `json:"choiceLabel,omitempty"`
		URL                         string   `json:"url,omitempty"`
		GenericID                   string   `json:"-"`
		Badges                      []Badge  `json:"badges,omitempty"`
		HasDigitallyGeneratedImages bool     `json:"hasDigitallyGeneratedImages,omitempty"`
	}

	type ProductList struct {
		ID                          string   `json:"id"` // The Product Page Instance id in CMS. Pattern: ^\\d{6}$
		GenericChoiceID             string   `json:"genericChoiceId"`
		URL                         string   `json:"url"`    // URL for this product on the site.
		Family                      string   `json:"family"` // Product family, e.g. 'Body by Victoria' (AKA 'brand').
		IsNew                       bool     `json:"isNew"`  // Flag that indicates that this is a new product.
		Name                        string   `json:"name"`   // Product name, e.g. 'Perfect Coverage Bra'. (AKA 'short description'.)
		Rating                      float64  `json:"rating"` // Average product rating, out of 5.
		TotalReviewCount            int      `json:"totalReviewCount"`
		GenericLabel                string   `json:"genericLabel,omitempty"` // "Label for this selector, e.g. 'Color' or 'Fragrance'."
		ProductImages               []string `json:"productImages"`          // Array of collection images.
		HasDigitallyGeneratedImages bool     `json:"hasDigitallyGeneratedImages,omitempty"`
		OfferCallout                []string `json:"offerCallout"`
		SwatchCount                 int64    `json:"swatchCount"`     // Number of available swatches for the product.
		SwatchLabel                 string   `json:"swatchLabel"`     // The label for product swatches (color, scent, team, etc.) from CMS
		IsZeroAttribute             bool     `json:"isZeroAttribute"` // Flag for whether a product can be purchased with only a quantity selection (1 size/1 choice).
		IsNewSavings                bool     `json:"isNewSavings"`    // Flag for whether to show new savings flag.
		IsGiftCard                  bool     `json:"isGiftCard"`      // Flag to indicate that this is a gift card product (but not an e-gift card).
		IsEGiftCard                 bool     `json:"isEGiftCard"`     // Flag to indicate that this is an e-gift card product (but not a gift card).
		IsClearance                 bool     `json:"isClearance"`
		ItemLevelCallout            []string `json:"itemLevelCallout"`
		Price                       string   `json:"price"`                      // Original price, formatted.
		AltPrices                   []string `json:"altPrices"`                  // An array of promotional or alternate prices for an item
		SalePrice                   string   `json:"salePrice"`                  // Sale price, formatted.
		CollectionShortDescription  string   `json:"collectionShortDescription"` // Collection short description
		HideOnMobile                bool     `json:"hideOnMobile"`               // Flag for whether product should be hidden on mobile devices
		DisplayGenericDescription   bool     `json:"displayGenericDescription"`
		Badges                      []Badge  `json:"badges,omitempty"`
		Swatches                    []Swatch `json:"swatches,omitempty"`
		Currency                    string   `json:"-"` // Placeholder for currency to help sort
		SortingPrice                float64  `json:"-"` // Placeholder for regular pricing to help sort
		Recommended                 bool     `json:"-"` // Placeholder for recommended products from personalization/certona
		MasterStyleId               string   `json:"masterStyleId"`
		CumulativeRating            float64  `json:"-"` // Placeholder for most rated to help sort
		Position                    int64    `json:"position,omitempty"`
		PromoID                     string   `json:"promoID,omitempty"`
		BundleType                  string   `json:"bundleType,omitempty"`
		UnitsSold                   float64  `json:"-"` // Placeholder for bestseller to help sort
	}

	// []Badge - 24 bytes
	// []string - 24 bytes
	// string - 16 bytes
	// int64 - 8 bytes
	// float64 - 8 bytes
	// bool - 1 byte
	type OptimisedProductList struct {
		Badges                      []Badge  `json:"badges,omitempty"`
		Swatches                    []Swatch `json:"swatches,omitempty"`
		ProductImages               []string `json:"productImages"` // Array of collection images.
		OfferCallout                []string `json:"offerCallout"`
		AltPrices                   []string `json:"altPrices"` // An array of promotional or alternate prices for an item
		ItemLevelCallout            []string `json:"itemLevelCallout"`
		ID                          string   `json:"id"` // The Product Page Instance id in CMS. Pattern: ^\\d{6}$
		MasterStyleId               string   `json:"masterStyleId"`
		GenericChoiceID             string   `json:"genericChoiceId"`
		URL                         string   `json:"url"`                        // URL for this product on the site.
		Family                      string   `json:"family"`                     // Product family, e.g. 'Body by Victoria' (AKA 'brand').
		Name                        string   `json:"name"`                       // Product name, e.g. 'Perfect Coverage Bra'. (AKA 'short description'.)
		Price                       string   `json:"price"`                      // Original price, formatted.
		Currency                    string   `json:"-"`                          // Placeholder for currency to help sort
		SalePrice                   string   `json:"salePrice"`                  // Sale price, formatted.
		CollectionShortDescription  string   `json:"collectionShortDescription"` // Collection short description
		SwatchLabel                 string   `json:"swatchLabel"`                // The label for product swatches (color, scent, team, etc.) from CMS
		PromoID                     string   `json:"promoID,omitempty"`
		BundleType                  string   `json:"bundleType,omitempty"`
		GenericLabel                string   `json:"genericLabel,omitempty"` // "Label for this selector, e.g. 'Color' or 'Fragrance'."
		Position                    int64    `json:"position,omitempty"`
		SwatchCount                 int64    `json:"swatchCount"` // Number of available swatches for the product.
		TotalReviewCount            int      `json:"totalReviewCount"`
		UnitsSold                   float64  `json:"-"`            // Placeholder for bestseller to help sort
		CumulativeRating            float64  `json:"-"`            // Placeholder for most rated to help sort
		SortingPrice                float64  `json:"-"`            // Placeholder for regular pricing to help sort
		Rating                      float64  `json:"rating"`       // Average product rating, out of 5.
		IsNew                       bool     `json:"isNew"`        // Flag that indicates that this is a new product.
		Recommended                 bool     `json:"-"`            // Placeholder for recommended products from personalization/certona
		HideOnMobile                bool     `json:"hideOnMobile"` // Flag for whether product should be hidden on mobile devices
		DisplayGenericDescription   bool     `json:"displayGenericDescription"`
		IsZeroAttribute             bool     `json:"isZeroAttribute"` // Flag for whether a product can be purchased with only a quantity selection (1 size/1 choice).
		IsNewSavings                bool     `json:"isNewSavings"`    // Flag for whether to show new savings flag.
		IsGiftCard                  bool     `json:"isGiftCard"`      // Flag to indicate that this is a gift card product (but not an e-gift card).
		IsEGiftCard                 bool     `json:"isEGiftCard"`     // Flag to indicate that this is an e-gift card product (but not a gift card).
		IsClearance                 bool     `json:"isClearance"`
		HasDigitallyGeneratedImages bool     `json:"hasDigitallyGeneratedImages,omitempty"`
	}

	// Struct Example
	// fmt.Println("Without structuring:", unsafe.Sizeof(Product{}))
	// fmt.Println("With structuring:", unsafe.Sizeof(OptimizedProduct{}))

	// Stacks follow up
	fmt.Println("Without structuring stacks:", unsafe.Sizeof(ProductList{}))
	fmt.Println("With structuring stacks:", unsafe.Sizeof(OptimisedProductList{}))
}
