const { sendNotification } = require("../src/notification_service");

describe("Mocking: Notification Service", () => {
  it("should return 'Notification Sent' when the send method is successful", () => {
    // Mock notificationService with a successful send method
    const notificationService = {
      send: jasmine.createSpy().and.returnValue(true), // Mocked send method returns true
    };

    const message = "Hello, World!";
    const result = sendNotification(notificationService, message);

    // Verify the send method is called with the correct message
    expect(notificationService.send).toHaveBeenCalledWith(message);

    // Verify the function returns "Notification Sent"
    expect(result).toBe("Notification Sent");
  });

  it("should return 'Failed to Send' when the send method fails", () => {
    // Mock notificationService with a failed send method
    const notificationService = {
      send: jasmine.createSpy().and.returnValue(false), // Mocked send method returns false
    };

    const message = "Hello, World!";
    const result = sendNotification(notificationService, message);

    // Verify the send method is called with the correct message
    expect(notificationService.send).toHaveBeenCalledWith(message);

    // Verify the function returns "Failed to Send"
    expect(result).toBe("Failed to Send");
  });
});
