package db

import (
	"errors"
	"github.com/cockroachdb/pebble"
)

type DB struct {
	db *pebble.DB
	wo *pebble.WriteOptions
}
type DataProcessor func(k, v []byte) (interface{}, error)

func NewDB(name string, options *pebble.Options) (*DB, error) {
	d, err := pebble.Open(name, options)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: d,
		wo: pebble.Sync,
	}, nil
}

func (d *DB) Get(k []byte, processor DataProcessor) (interface{}, error) {
	bytes, closer, err := d.db.Get(k)
	if err != nil {
		return nil, err
	}
	err = closer.Close()
	if err != nil {
		return nil, err
	}
	result, err := processor(nil, bytes)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DB) BatchGet(ks [][]byte, processor DataProcessor) ([]interface{}, error) {
	batch := d.db.NewBatch()
	result := make([]interface{}, len(ks))
	for i, k := range ks {
		v, closer, err := batch.Get(k)
		if closer != nil {
			err = closer.Close()
			if err != nil {
				return nil, err
			}
		}
		if err != nil {
			return nil, err
		}
		r, err := processor(k, v)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}

func (d *DB) Iter(start []byte, end []byte, toLast bool, processor DataProcessor) ([]interface{}, error) {
	iterOptions := &pebble.IterOptions{
		LowerBound: start,
	}
	if !toLast {
		iterOptions.UpperBound = end
	}
	var result []interface{}
	iter := d.db.NewIter(iterOptions)
	if iter.First() {
		for {
			k := iter.Key()
			v := iter.Value()
			r, err := processor(k, v)
			if err != nil {
				return nil, err
			}
			if r != nil {
				result = append(result, r)
			}
			if !iter.Next() {
				break
			}
		}
	}
	return result, nil
}

func (d *DB) Set(k, v []byte) error {
	return d.db.Set(k, v, d.wo)
}

func (d *DB) BatchSet(ks, vs [][]byte) error {
	if len(ks) != len(vs) {
		return errors.New("batch keys与values数量不匹配")
	}
	batch := d.db.NewBatch()
	for i := 0; i < len(ks); i++ {
		k := ks[i]
		v := vs[i]
		err := batch.Set(k, v, d.wo)
		if err != nil {
			_ = batch.Close()
			return err
		}
	}
	return batch.Commit(d.wo)
}

func (d *DB) Last(processor DataProcessor) ([]byte, []byte, error) {
	iter := d.db.NewIter(&pebble.IterOptions{})
	if iter.Last() {
		k := iter.Key()
		v := iter.Value()
		_, err := processor(k, v)
		return k, v, err
	}
	return nil, nil, nil
}
