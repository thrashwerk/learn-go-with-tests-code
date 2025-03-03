package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})

	// AssertEqual(t, 1, "1") // uncomment to see the error
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(StackOfInts)

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(456)
		v, _ := myStackOfInts.Pop()
		AssertEqual(t, v, 456)

		v, _ = myStackOfInts.Pop()
		AssertEqual(t, v, 123)

		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(StackOfStrings)

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("456")
		v, _ := myStackOfStrings.Pop()
		AssertEqual(t, v, "456")

		v, _ = myStackOfStrings.Pop()
		AssertEqual(t, v, "123")

		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}
