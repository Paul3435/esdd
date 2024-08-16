# How to use

## Prerequisites
Ensure you have the following installed before proceeding:
- [Docker](https://docs.docker.com/get-docker/)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Instructions
- Clone the repository in your environment:
`git clone https://github.com/Paul3435/esdd.git`

*It may be necessary to specify a [Personal Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) in the Password field.*


- Navigate to the new folder:
`cd esdd`

- Build a docker image using the Dockerfile as blueprint:
`docker build -t dreamdatapp .`

- Run the image in a container:
`docker run -it -p 8080:8080 dreamdatapp`

- You are given two choices now:
    - Select **Test** mode, which is a terminal-based application that showcases the methods for sending emails and the possibility to build custom emails and keys.

    - Skip (press Enter or enter any string other than 'test'), which will then host the API in localhost:8080.

- The API contains 2 endpoint possibilities:
    - **/**: leads to the [landing page](https://imgur.com/a/0AFHldu)
    - **/send**: leads to the [email form](https://imgur.com/a/763lhMN). Filling out the form and sending an email will debug to the terminal exclusively.
    (Mailgun only allows for sandbox domains with verified receivers (paul.borjesson.sesma3435@gmail.com), therefore emails with any other target will fail)

---

# Quick description of my approach:

The solution is a resilient email service that integrates with both SendGrid and Mailgun. The reason for choosing these two specific providers is that they didn't require Custom Domains. SendGrid allows for [Single Sender Verification](https://www.twilio.com/docs/sendgrid/ui/sending-email/sender-verification), and Mailgun allows [Sandbox Domains](https://help.mailgun.com/hc/en-us/articles/217531258-Authorized-Recipients).

The core logic is encapsulated in an EmailService interface, which is implemented by each email provider, thus allowing the possibility to add more providers in an intuitive and schematic manner. The EmailServiceManager manages these providers, attempting to send emails via the primary service and falling back to secondary services upon failure, this is done by accepting by parameter a slice of providers and cycling through them.

Retry logic has not been implemented, as while transient issues may occur, it seems redundant to add this feature when fallbacks are already present. However, should I be mistaken, taking into consideration that this feature should only require a single line of code in pkg/email/email.go's SendEmail method, feel free to suggest changing it! :)

An HTTP server is provided for basic user interaction, allowing emails to be sent via a web form. This also opens up the possibility of implementing new features in an intuitive manner by adding more handlers to new endpoints.

Additionally, a command-line test mode has been implemented in order to facilitate testing without the need for a running server. It also adds flexibility to the request formations (such as using your own ApiKey), allowing reviewers to verify the service's behavior in various scenarios.

To facilitate deployment and testing, a Dockerfile is included, allowing the service to be containerized and run consistently across environments.

# Known issues:
- Loading the email form will send a POST request, which will fail as the fields aren't filled.

# To-do features:
- Implement automatic unit testing.
- Improve the UI in order to provide more feedback to the user.
- Set up API rate limits.
- Set up a Load Balancer in order to use all the providers in an equal manner. 
- (Admin) Set up monitoring tools to track the usage delivery rates, provider performance, etc.
- Add more providers in order to improve redundancy and increase the fallback options.
- Better handling of the API KEYs.
