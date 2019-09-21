// Package tdigest provides a simple and (memory) efficient way to
// compute distribution quartiles on the fly from a potentially
// large number of data points.
//
// It is (freely) inspired by the paper from T.Dunning :
// https://github.com/tdunning/t-digest/blob/master/docs/t-digest-paper/histo.pdf
//
// As new data points are added, the key parameter is the choice of the Sizer,
// that determines how aggressively the buckets used to aggregate the data are merged.
// The merging process happens regularly
// when data points are added, and before computing anything.
// A map-reduce approach is also achievable, since  TD structures can be computed
// in parallel and then merged. When merging, both Sizer are expected to be identical.
//
package tdigest
