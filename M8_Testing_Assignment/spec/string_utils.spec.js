const { capitalize, reverseString } = require("../src/string_utils");

describe("String Utility Functions", () => {
  // Tests for the capitalize function
  describe("capitalize", () => {
    it("should capitalize the first letter of a word", () => {
      expect(capitalize("hello")).toBe("Hello");
    });

    it("should return an empty string if input is empty", () => {
      expect(capitalize("")).toBe("");
    });

    it("should capitalize single-character words", () => {
      expect(capitalize("a")).toBe("A");
    });

    it("should handle strings that are already capitalized", () => {
      expect(capitalize("Hello")).toBe("Hello");
    });

    it("should not modify the rest of the string", () => {
      expect(capitalize("hELLO")).toBe("HELLO");
    });
  });

  // Tests for the reverseString function
  describe("reverseString", () => {
    it("should reverse a string", () => {
      expect(reverseString("hello")).toBe("olleh");
    });

    it("should handle empty strings", () => {
      expect(reverseString("")).toBe("");
    });

    it("should handle palindromes correctly", () => {
      expect(reverseString("madam")).toBe("madam");
    });

    it("should reverse a single-character string", () => {
      expect(reverseString("a")).toBe("a");
    });

    it("should reverse a string with spaces", () => {
      expect(reverseString("hello world")).toBe("dlrow olleh");
    });
  });
});
