const { toggleVisibility } = require("../src/dom_manipulation");
describe("Spying: DOM Manipulation - toggleVisibility", () => {
  let element;

  beforeEach(() => {
    // Create a mock DOM element with a spy for 'style.display'
    element = {
      style: {
        _display: "block", // Internal mock property to hold the display value
        get display() {
          return this._display;
        },
        set display(value) {
          this._display = value;
        },
      },
    };

    spyOnProperty(element.style, "display", "set").and.callThrough();
    spyOnProperty(element.style, "display", "get").and.callThrough();
  });

  it("should set display to 'none' if the element is initially visible", () => {
    toggleVisibility(element);

    expect(element.style.display).toBe("none");
    expect(element.style.display).toBe("none");
  });

  it("should set display to 'block' if the element is initially hidden", () => {
    // Set the initial display to 'none'
    element.style.display = "none";

    toggleVisibility(element);

    expect(element.style.display).toBe("block");
    expect(element.style.display).toBe("block");
  });
});
