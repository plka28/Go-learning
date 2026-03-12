package bins

import "time"

type binList = []Bin

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}, nil
}
