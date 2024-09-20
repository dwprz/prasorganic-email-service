
![Logo](https://ik.imagekit.io/pj3r6oe9k/prasorganic-high-resolution-logo-transparent.svg?updatedAt=1726835541390)

# Prasorganic Email Service

Prasorganic Email Service is one of the components in the Prasorganic architecture Microservice built with Go (Golang), this service functions as a consumer RabbitMQ and email sending. RabbitMQ was chosen as a message broker because email sending needs that do not require high throughput, so it is suitable with simple application requirements.

## Tech Stack

[![My Skills](https://skillicons.dev/icons?i=go,rabbitmq,bash,git&theme=light)](https://skillicons.dev)

## Features

- **OTP Delivery:** Implements OTP delivery via Gmail using the Google API Gmail v1, configured with OAuth2 to enhance security by eliminating the need to send credentials, thereby reducing the risk of credential leakage.

- **Message Broker:** Consumes messages from RabbitMQ for email delivery.

- **Logging:** Logs are recorded using Logrus.

- **Error Handling:** Implements error handling to ensure proper detection and handling of errors, minimizing the impact on both the client and server.

- **Configuration and Security:** Utilizes Viper and HashiCorp Vault for integrated configuration and security management.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

This project makes use of third-party packages and tools. The licenses for these
dependencies can be found in the `LICENSES` directory.

## Dependencies and Their Licenses

- `Go:` Licensed under the BSD 3-Clause "New" or "Revised" License. For more information, see the [Go License](https://github.com/golang/go/blob/master/LICENSE).

- `GNU Make:` Licensed under the GNU General Public License v3.0. For more information, see the [GNU Make License](https://www.gnu.org/licenses/gpl.html).

- `GNU Bash:` Licensed under the GNU General Public License v3.0. For more information, see the [Bash License](https://www.gnu.org/licenses/gpl-3.0.html).

- `Git:` Licensed under the GNU General Public License version 2.0. For more information, see the [Git License](https://opensource.org/license/GPL-2.0).

## Thanks 👍
Thank you for viewing my project.