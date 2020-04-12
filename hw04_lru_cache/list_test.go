package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func checkEmptyList(t *testing.T, l List) {
	require.Equal(t, l.Len(), 0)
	require.Nil(t, l.Front())
	require.Nil(t, l.Back())
}

func TestList(t *testing.T) {

	t.Run("empty list", func(t *testing.T) {
		checkEmptyList(t, NewList())
	})

	t.Run("push_front", func(t *testing.T) {
		l := NewList()

		firstVal := 5
		l.PushFront(firstVal) // [5]
		require.NotNil(t, l.Front())
		require.Equal(t, firstVal, l.Front().Value)
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Front().Prev)
		require.Equal(t, l.Front(), l.Back())

		secondVal := "some test string"
		l.PushFront(secondVal) // ["some test string", 5]
		require.NotNil(t, l.Front())
		require.Equal(t, secondVal, l.Front().Value)
		require.Nil(t, l.Front().Next)
		require.NotNil(t, l.Front().Prev)
		require.Equal(t, firstVal, l.Front().Prev.Value)
		require.Nil(t, l.Front().Prev.Prev)

		thirdVal := 2
		l.PushFront(thirdVal) // [2, "some test string", 5]
		require.NotNil(t, l.Front())
		require.Equal(t, thirdVal, l.Front().Value)
		require.Nil(t, l.Front().Next)
		require.NotNil(t, l.Front().Prev)
		require.Equal(t, secondVal, l.Front().Prev.Value)
		require.NotNil(t, l.Front().Prev.Prev)
		require.Equal(t, firstVal, l.Front().Prev.Prev.Value)

		require.Equal(t, firstVal, l.Back().Value)
	})

	t.Run("push_back", func(t *testing.T) {
		l := NewList()

		firstVal := 5
		l.PushBack(firstVal) // [5]
		require.NotNil(t, l.Back())
		require.Equal(t, 5, l.Back().Value)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Back().Prev)
		require.Equal(t, l.Front(), l.Back())

		secondVal := "some test string"
		l.PushBack(secondVal) // [5, "some test string"]
		require.NotNil(t, l.Back())
		require.Equal(t, secondVal, l.Back().Value)
		require.Nil(t, l.Back().Prev)
		require.NotNil(t, l.Back().Next)
		require.Equal(t, firstVal, l.Back().Next.Value)
		require.Nil(t, l.Back().Next.Next)

		thirdVal := 2
		l.PushBack(thirdVal) // [5, "some test string", 2]
		require.NotNil(t, l.Back())
		require.Equal(t, thirdVal, l.Back().Value)
		require.Nil(t, l.Back().Prev)
		require.NotNil(t, l.Back().Next)
		require.Equal(t, secondVal, l.Back().Next.Value)
		require.NotNil(t, l.Back().Next.Next)
		require.Equal(t, firstVal, l.Back().Next.Next.Value)

		require.Equal(t, firstVal, l.Front().Value)
	})

	t.Run("remove", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Back().Next // 20
		l.Remove(middle)        // [10, 30]
		require.Equal(t, 2, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Front().Prev.Value)
		require.Equal(t, 10, l.Back().Next.Value)
		require.Equal(t, 30, l.Back().Value)

		l.Remove(l.Back())
		require.Equal(t, 1, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)

		l.Remove(l.Front())
		checkEmptyList(t, l)
	})

	t.Run("moveToFront", func(t *testing.T) {
		l := NewList()
		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]

		l.MoveToFront(l.Front()) // [10, 20, 30]
		require.Equal(t, 10, l.Front().Value)

		l.MoveToFront(l.Back()) // [30, 10, 20]
		require.Equal(t, 30, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Back().Next // 20
		l.Remove(middle)        // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Back(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, elems)
	})
}
