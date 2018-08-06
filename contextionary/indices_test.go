/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
package contextionary

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/creativesoftwarefdn/weaviate/contextionary/generator"
)

// Test data
// Changing this data might invalidate the nearest neighbours test!
var vectorTests = []struct {
	word string
	vec  []float32
}{
	{"apple", []float32{1, 0, 0}},
	{"pie", []float32{0, 1, 0}},
	{"computer", []float32{0, 0, 1}},
	{"fruit", []float32{0.8, 0, 0}},
	{"company", []float32{0, 0, 2}},
}

func TestMMappedIndex(t *testing.T) {
	tempdir, err := ioutil.TempDir("", "weaviate-vector-test")

	if err != nil {
		t.Errorf("Could not create temporary directory, %v", err)
	}

	defer os.RemoveAll(tempdir)

	// First generate the csv input fileformat based on the test data.
	var dataset = ""

	for i := 0; i < len(vectorTests); i++ {
		vt := vectorTests[i]
		dataset += vt.word + " "
		for j := 0; j < len(vt.vec)-1; j++ {
			dataset += fmt.Sprintf("%f ", vt.vec[j])
		}
		dataset += fmt.Sprintf("%f\n", vt.vec[len(vt.vec)-1])
	}

	err = ioutil.WriteFile(tempdir+"/glove.txt", []byte(dataset), 0644)
	if err != nil {
		t.Errorf("Could not create input file: %v", err)
	}

	t.Run("Generating index", func(t *testing.T) {
		// Now build an index based on this
		var genOpts generator.Options
		genOpts.VectorCSVPath = tempdir + "/glove.txt"
		genOpts.TempDBPath = tempdir + "/tempdb"
		genOpts.OutputPrefix = tempdir + "/glove"
		genOpts.K = 3
		generator.Generate(genOpts)
	})

	// And load the index.
	vi, err := LoadVectorFromDisk(tempdir+"/glove.knn", tempdir+"/glove.idx")
	if err != nil {
		t.Errorf("Could not load vectors from disk: %v", err)
	}

	sharedTests(t, vi)
}

func TestInMemoryIndex(t *testing.T) {
	builder := InMemoryBuilder(3)
	for i := 0; i < len(vectorTests); i++ {
		v := vectorTests[i]
		builder.AddWord(v.word, NewVector(v.vec))
	}

	memoryIndex := Contextionary(builder.Build(3))

	sharedTests(t, &memoryIndex)
}

func TestCombinedIndex(t *testing.T) {
	builder1 := InMemoryBuilder(3)
	builder2 := InMemoryBuilder(3)

	split := 3

	for i := 0; i < split; i++ {
		v := vectorTests[i]
		builder1.AddWord(v.word, NewVector(v.vec))
	}

	for i := split; i < len(vectorTests); i++ {
		v := vectorTests[i]
		builder2.AddWord(v.word, NewVector(v.vec))
	}

	memoryIndex1 := Contextionary(builder1.Build(3))
	memoryIndex2 := Contextionary(builder2.Build(3))

	indices12 := []Contextionary{memoryIndex1, memoryIndex2}
	indices21 := []Contextionary{memoryIndex2, memoryIndex1}

	t.Run("indices 1,2", func(t *testing.T) { testCombined(t, indices12) })
	t.Run("indices 2,1", func(t *testing.T) { testCombined(t, indices21) })
}

func testCombined(t *testing.T, indices []Contextionary) {
	combinedIndex, err := CombineVectorIndices(indices)
	if err != nil {
		t.Errorf("Combining failed")
		t.FailNow()
	}

	err = combinedIndex.VerifyDisjoint()

	if err != nil {
		t.Errorf("Not disjoint; %v", err)
		t.FailNow()
	}

	vi := Contextionary(combinedIndex)

	sharedTests(t, &vi)
}

func sharedTests(t *testing.T, vi *Contextionary) {
	t.Run("Number of elements is correct", func(t *testing.T) {
		expected := 5
		found := (*vi).GetNumberOfItems()
		if found != expected {
			t.Errorf("Expected to have %v items, but found %v", expected, found)
		}
	})

	t.Run("Iterate over all items", func(t *testing.T) {
		// Iterate over all items. Check index -> word, and lookup word -> index
		length := ItemIndex((*vi).GetNumberOfItems())
		for i := ItemIndex(0); i < length; i++ {
			word, err := (*vi).ItemIndexToWord(ItemIndex(i))
			if err != nil {
				t.Errorf("Could not get item of index %+v, because: %+v", i, err)
			}

			i2 := (*vi).WordToItemIndex(word)

			if i2 != i {
				t.Errorf("Index -> Word -> Index failed!. i=%v, w=%v i2=%v", i, word, i2)
			}

		}
	})

	t.Run("Check that feature vectors are stored properly", func(t *testing.T) {
		for i := 0; i < len(vectorTests); i++ {
			vt := vectorTests[i]
			wordIndex := (*vi).WordToItemIndex(vt.word)
			if !wordIndex.IsPresent() {
				t.Errorf("Could not find word %v", vt.word)
			}
			// Get back the feature vectors.
			vector, err := (*vi).GetVectorForItemIndex(wordIndex)
			if err != nil {
				t.Errorf("Could not get vector")
			}

			if vector == nil {
				t.Errorf("Vector missing!")
				t.FailNow()
			}

			// and check that it's correct
			vtvec := NewVector(vt.vec)
			areEqual, err := vector.Equal(&vtvec)
			if err != nil {
				t.Errorf("Could not compare the two vectors: %v", err)
			}

			if !areEqual {
				t.Errorf("Feature vector %v incorrect (word: %v). Expected %v, got %v", i, vt.word, vt.vec, vector.vector)
			}
		}
	})

	t.Run("Test that the distances between all pairs of test data is correct", func(t *testing.T) {
		for i := 0; i < len(vectorTests); i++ {
			for j := 0; j < len(vectorTests); j++ {
				vtA := vectorTests[i]
				vtB := vectorTests[j]
				vtAVec := NewVector(vtA.vec)
				vtBVec := NewVector(vtB.vec)

				wiA := (*vi).WordToItemIndex(vtA.word)
				wiB := (*vi).WordToItemIndex(vtB.word)

				annoyDist, err := (*vi).GetDistance(wiA, wiB)
				if err != nil {
					t.Errorf("Could not compute distance")
				}

				simpleDist, err := vtAVec.Distance(&vtBVec)
				if err != nil {
					panic("should be same length")
				}

				if !equalFloatEpsilon(annoyDist, simpleDist, 0.00003) {
					t.Errorf("Distance between %v and %v incorrect; %v (annoy) vs %v (test impl)", vtA.word, vtB.word, annoyDist, simpleDist)
				}
			}
		}
	})

	t.Run("Test nearest neighbours apple & fruit", func(t *testing.T) {
		appleIdx := (*vi).WordToItemIndex("apple")
		fruitIdx := (*vi).WordToItemIndex("fruit")

		res, distances, err := (*vi).GetNnsByItem(fruitIdx, 2, 3)
		t.Logf("%v, %v, %v\n", res, distances, err)
		if err != nil {
			t.Errorf("GetNNs failed!")
		}
		if len(res) != 2 {
			t.Errorf("Wrong number of items returned; got %v expected 2", len(res))
		}
		// res[0] will be fruit itself.
		if res[0] != fruitIdx {
			closestTo, _ := (*vi).ItemIndexToWord(res[0])
			t.Errorf("closest element should be itself, fruit, but was '%v'. all results:\n%v", closestTo, debugPrintItems(vi, res, distances))
		}

		if res[1] != appleIdx {
			closestTo, _ := (*vi).ItemIndexToWord(res[1])
			t.Errorf("apple should be closest to apple, but was '%v'. all results:\n%v", closestTo, debugPrintItems(vi, res, distances))
		}

		if !equalFloatEpsilon(distances[1], 0.2, 0.0002) {
			t.Errorf("Wrong distances!, got %v", distances[1])
		}
	})

	t.Run("Test nearest neighbours computer & company", func(t *testing.T) {
		companyIdx := (*vi).WordToItemIndex("company")
		computerIdx := (*vi).WordToItemIndex("computer")

		res, distances, err := (*vi).GetNnsByItem(companyIdx, 2, 3)
		if err != nil {
			t.Errorf("GetNNs failed!")
		}
		if len(res) != 2 {
			t.Errorf("Wrong number of items returned; got %v expected 2", len(res))
			t.FailNow()
		}
		// res[0] will be company itself.
		if res[1] != computerIdx {
			t.Errorf("computer should be closest to company!")
		}

		if !equalFloatEpsilon(distances[1], 1, 0.0002) {
			t.Errorf("Wrong distances!, got %v", distances[1])
		}
	})

	t.Run("Test k-nearest from vector", func(t *testing.T) {
		var applePie = NewVector( /* centroid of apple and pie */ []float32{0.5, 0.5, 0})

		fruitIdx := (*vi).WordToItemIndex("fruit")
		appleIdx := (*vi).WordToItemIndex("apple")
		pieIdx := (*vi).WordToItemIndex("pie")

		res, distances, err := (*vi).GetNnsByVector(applePie, 3, 3)
		if err != nil {
			t.Errorf("GetNNs failed: %v", err)
			t.FailNow()
		}
		if len(res) != 3 {
			t.Errorf("Wrong number of items returned; got %v expected 3", len(res))
			t.FailNow()
		}

		if res[0] != fruitIdx {
			closestTo, _ := (*vi).ItemIndexToWord(res[1])
			t.Errorf("fruit should be closest to fruit!, but was '%v'", closestTo)
			t.Errorf("got results: %+v", res)
			for _, i := range res {
				word, err := (*vi).ItemIndexToWord(i)
				t.Errorf("got word: %v (err: %v)", word, err)
			}
		}

		if res[1] != appleIdx {
			closestTo, _ := (*vi).ItemIndexToWord(res[1])
			t.Errorf("apple should be 2nd closest to apple!, but was '%v'", closestTo)
		}

		if res[2] != pieIdx {
			closestTo, _ := (*vi).ItemIndexToWord(res[2])
			t.Errorf("pie should be 3rd closest to pie!, but was '%v'", closestTo)
		}

		vFruit, err := (*vi).GetVectorForItemIndex(fruitIdx)
		if err != nil {
			t.Errorf("could not fetch fruit vector")
			return
		}

		vApple, err := (*vi).GetVectorForItemIndex(appleIdx)
		if err != nil {
			panic("could not fetch apple vector")
		}

		vPie, err := (*vi).GetVectorForItemIndex(pieIdx)
		if err != nil {
			panic("could not fetch pie vector")
		}

		distanceToFruit, err := applePie.Distance(vFruit)
		if err != nil {
			panic("should be same length")
		}
		if !equalFloatEpsilon(distances[0], distanceToFruit, 0.0001) {
			t.Errorf("Wrong distance for fruit, expect %v, got %v", distanceToFruit, distances[0])
		}

		distanceToApple, err := applePie.Distance(vApple)
		if err != nil {
			panic("should be same length")
		}
		if !equalFloatEpsilon(distances[1], distanceToApple, 0.0001) {
			t.Errorf("Wrong distance for apple, got %v", distances[1])
		}

		distanceToPie, err := applePie.Distance(vPie)
		if err != nil {
			panic("should be same size")
		}
		if !equalFloatEpsilon(distances[2], distanceToPie, 0.0001) {
			t.Errorf("Wrong distance for pie, expected %v, got %v", distanceToPie, distances[2])
		}
	})
}

func equalFloatEpsilon(a float32, b float32, epsilon float32) bool {
	var min, max float32

	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}

	return max <= (min + epsilon)
}

func debugPrintItems(vi *Contextionary, items []ItemIndex, distances []float32) string {
	result := ""
	for i, item := range items {
		w, _ := (*vi).ItemIndexToWord(item)
		result += fmt.Sprintf("%v: %v (%v)\n", item, w, distances[i])
	}

	return result
}
