# GitHub User Activity

This project is a backend component for the Roadmap.sh project, specifically designed to retrieve and display recent activity for a given GitHub user.  It's categorized as a beginner-level project and is part of the backend learning path.

[View this project on Roadmap.sh](https://roadmap.sh/projects/github-user-activity)

## Functionality

This Go application fetches recent GitHub events for a specified user using the GitHub API.  It then processes and formats these events into a user-friendly output, displaying information such as:

* **Push events:** Number of commits pushed to a repository.
* **Issues events:** Actions performed on issues (e.g., opened, closed, edited).
* **Issue comment events:** Indication of new comments on issues.
* **Watch events:**  Indicates when a user starred a repository.
* **Other event types:**  Generic display of other event types.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

* Go 1.18 or higher installed on your system.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/[your-github-username]/github-activity.git

2. Navigate to the project directory:
    ```bash
    cd github-activity

### Usage
1.  Build the application:

    ```bash
    go build github-activity.go
    ```

2. Run the application, providing the GitHub username as a command-line argument:

    ```bash
    ./github-activity [GitHub_username]
    ```

    For example:

    ```bash
    ./github-activity alielmi98
    ```

This will print the recent activity for the specified GitHub user to your console.

### Built With
Go - Programming language

### Contributing
Contributions are welcome! Please feel free to open issues or submit pull requests.

### License
This project is licensed under the MIT License - see the LICENSE file for details.
