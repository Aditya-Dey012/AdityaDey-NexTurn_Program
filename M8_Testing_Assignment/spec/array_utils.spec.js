const { getElement } = require("../src/array_utils");

describe("Array Utilities - getElement", () => {
  // Tests for valid index values
  it("should return the element at a valid index", () => {
    const arr = [10, 20, 30, 40, 50];
    expect(getElement(arr, 0)).toBe(10); // First element
    expect(getElement(arr, 2)).toBe(30); // Middle element
    expect(getElement(arr, 4)).toBe(50); // Last element
  });

  // Test for negative index
  it("should throw an error for negative indices", () => {
    const arr = [10, 20, 30];
    expect(() => getElement(arr, -1)).toThrowError("Index out of bounds");
  });

  // Test for out-of-range index
  it("should throw an error for indices greater than or equal to array length", () => {
    const arr = [10, 20, 30];
    expect(() => getElement(arr, 3)).toThrowError("Index out of bounds"); // Array length is 3, so index 3 is out of bounds
    expect(() => getElement(arr, 100)).toThrowError("Index out of bounds"); // Large out-of-range index
  });

  // Test for empty arrays
  it("should throw an error when accessing any index of an empty array", () => {
    const arr = [];
    expect(() => getElement(arr, 0)).toThrowError("Index out of bounds");
    expect(() => getElement(arr, -1)).toThrowError("Index out of bounds");
  });

  // Test for valid indices in a single-element array
  it("should return the element if the index is valid in a single-element array", () => {
    const arr = [42];
    expect(getElement(arr, 0)).toBe(42);
  });
});
