# Golang

## Reminder Scheduling System

Create a scheduler system using at least two Golang services. The services will work as the backend for an alarm system in which users can create alarms/reminders and share them with other users.

User should be able to create new schduled alarms with just a datetime, a small description and a list of users to send the alarm to.

Acceptance Criteria:
- Create/deploy the backend services needed to support the system described above 
- Incoming  requests will provide only user ids 
- Data should flow in the following manner
  - Upon receiving a schedule request, we need to:
    - Validate that the request data is in correct format
    - Validate the user ids listed in the request
    - If users are valid, we need to create a new reminder entry in the database
    - We also need to start tracking the time until it's time to raise the alarms, use `goroutines` to do this concurrently
  - Once it is time to trigger an alarm, we should:
    - Retrieve the reminder entry from the database
    - Retrieve the email addresses for the users listed in the reminder entry
    - "Send" the reminder entry with complete info to the output table
- Organize the project code into meaningful packages, look at [intro3](https://github.com/vorticist/golang-intro/blob/main/intro3/intro3.md#packages) for a suggested structure
- Two main services
  - One will be a data API, a simple API app that will retrieve user data from a postgres database.
  - Scheduler, a service that will receive schedule request and will start tracking scheduled alarms
- Database
  - Use a separate schema for users
  - Use a separate schema for scheduled items
  - Use a separate schema for output

| users   | type  |
| :------ | :---: |
| user_id | uuid  |
| email   | text  |


| scheduled_items |  type  |
| :-------------- | :----: |
| id              |  uuid  |
| description     |  text  |
| users           | uuid[] |

| output      |  type  |
| :---------- | :----: |
| id          |  uuid  |
| description |  text  |
| emalis      | text[] |
