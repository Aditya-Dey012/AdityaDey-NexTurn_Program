function capitalize(word) {
    if (!word) return ""; // Handle empty input
    return word[0].toUpperCase() + word.slice(1); // Capitalize the first letter
  }
  
  function reverseString(str) {
    return str.split("").reverse().join(""); // Reverse the string
  }
  
  // Export the functions for testing
  module.exports = { capitalize, reverseString };
  