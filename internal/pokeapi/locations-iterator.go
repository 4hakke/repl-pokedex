package pokeapi

type LocationsIterator struct {
	Limit  int
	offset int
	count  *int
}

func (li *LocationsIterator) Next() ([]Location, error) {
	if li.count != nil && li.offset >= *li.count {
		return []Location{}, nil
	}
	res, err := locations(li.offset, li.Limit)
	if err != nil {
		return []Location{}, err
	}
	li.offset += li.Limit
	li.count = &res.Count
	return res.Results, nil
}

func (li *LocationsIterator) Previous() ([]Location, error) {
	if li.offset-li.Limit*2 < 0 {
		return []Location{}, nil
	}
	li.offset -= li.Limit * 2
	res, err := locations(li.offset, li.Limit)
	if err != nil {
		return []Location{}, err
	}
	li.count = &res.Count
	li.offset += li.Limit
	return res.Results, nil
}
