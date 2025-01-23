const { delayedGreeting } = require("../src/async_utils");

describe("Async Utilities - delayedGreeting", () => {
    beforeEach(() => {
        jasmine.clock().install(); // Install Jasmine's clock
    });

    afterEach(() => {
        jasmine.clock().uninstall(); // Uninstall Jasmine's clock
    });

    it("should resolve with the correct greeting message", async () => {
        const promise = delayedGreeting("John", 1000); // Start the promise
        jasmine.clock().tick(1000); // Fast-forward the timer by 1000ms

        await expectAsync(promise).toBeResolvedTo("Hello, John!"); // Assert the resolved value
    });

    it("should respect the specified delay before resolving", async () => {
        const promise = delayedGreeting("Jane", 2000); // Start the promise
        jasmine.clock().tick(2000); // Fast-forward the timer by 2000ms

        await expectAsync(promise).toBeResolvedTo("Hello, Jane!"); // Assert the resolved value
    });
});
