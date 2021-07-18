package filters

func BuildProductsFilters(productType *string, minPriceInCents, maxPriceInCents *int) []FilterPair {

	filters := make([]FilterPair, 0)

	if productType != nil {
		var typeFilter = FilterPair{ConditionString: "type = ?", Values: []interface{}{productType}}
		filters = append(filters, typeFilter)
	}

	var priceFilter FilterPair
	if minPriceInCents != nil {
		if maxPriceInCents != nil {
			priceFilter = FilterPair{
				ConditionString: "price_in_cents BETWEEN ? AND ?",
				Values:          []interface{}{*minPriceInCents, *maxPriceInCents}}
		} else {
			priceFilter = FilterPair{
				ConditionString: "price_in_cents >= ?",
				Values:          []interface{}{*minPriceInCents}}
		}
	} else {
		if maxPriceInCents != nil {
			priceFilter = FilterPair{
				ConditionString: "price_in_cents <= ?",
				Values:          []interface{}{*maxPriceInCents}}
		}
	}

	filters = append(filters, priceFilter)

	return filters
}
