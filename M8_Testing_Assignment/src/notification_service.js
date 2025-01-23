function sendNotification(notification_service, message) {
    const status = notification_service.send(message);
    return status ? "Notification Sent" : "Failed to Send";
  }
  
  module.exports = { sendNotification };
  