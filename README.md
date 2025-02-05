# dev-ops

## Porsche API

<img src="https://logodownload.org/wp-content/uploads/2021/02/porsche-logo-0.png"/>

Free fake API, written with <a href="https://go.dev/" target="_blank">Golang</a> for testing and prototyping.

### Why ?
Are you tired of spending valuable time registering for complex APIs when all you need is some data for testing or prototyping?

Introducing Porsche API – a lightweight Golang-based API, designed to provide quick and easy access to dummy data without the hassle of registration or complex API documentation.

### Features
* No registration
* Basic API
* Cross-domain
* Supports GET, POST, PUT, DELETE
* HTTP
* Compatible with React, Angular, Vue, Ember, ...

### Set up locally

Before you dive in, make sure to set up your environment variables by following these steps:

1. Clone the project: Open your terminal and clone our project repository by running the following command:
`git clone https://github.com/vitolinho/dev-ops.git`

2. **Grab the `.env.example`**: Head over to your project directory and find the `.env.example` file. This serves as a template for your environment variables.

3. **Fill in the `.env` file**: Duplicate the `.env.example` file and rename it to `.env`. Open it up and fill in the required values based on your setup. Don't worry, the example file provides hints on what each variable should be set to.

4. **Ready to Launch!**: Once your `.env` file is properly filled, you're almost ready to go!

5. **Start Docker**: If you haven't already, make sure you have Docker installed. [Install Docker](https://docs.docker.com/get-docker/) if you haven't already.

6. **Run the command**: In your terminal, simply execute `make up` or `docker-compose up -d --build` to start up your local environment.


### Terraform

To provision infrastructure using Terraform, follow these steps:

1. **Configure your Terraform variables**:
   Ensure that the required values are set in `terraform/terraform.tfvars`:
   ```hcl
   aws_access_key = "YOUR_AWS_ACCESS_KEY"
   aws_secret_key = "YOUR_AWS_SECRET"
   ssh_user = "YOUR_SSH_USER"

2. **Configure your GitHub token**:  
   Ensure that your personal access token is set in `terraform/ansible/vars.yml`:  
   ```yaml
   github_token: "YOUR_GITHUB_TOKEN"
   ```  
   You can generate a new GitHub token by visiting [GitHub Tokens](https://github.com/settings/tokens), clicking on **Generate new token (classic)**, giving it a name, and selecting the **admin:public_key** option.

3. **Initialize Terraform**:  
   `terraform init`

4. **Plan your changes** (optional but recommended to preview changes):  
   `terraform plan`

5. **Apply changes to provision infrastructure**:  
   `terraform apply -auto-approve`

6. **Destroy infrastructure** (if needed):  
   `terraform destroy -auto-approve`

7. **SSH into a newly created server**:  
   If you wish to connect to one of the created servers via SSH, automatically generated `make` commands are available. For example:  
   `make ssh-dev-germany-instance`  
   This command allows you to quickly access the server without manually specifying SSH parameters.

### Ansible
Ansible is automatically triggered by Terraform during the infrastructure provisioning process. However, if modifications are made to the Ansible configuration, you can run it manually. Below is an example command to execute Ansible manually:

```bash
ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook \
   -i ansible/inventory/prod-germany-instance \
   -u {ssh_user} \
   --private-key keys/prod-germany-instance-key.pem \
   ansible/playbook.yml
```

### Linter
First install the linter with this command
```bash
cd api/scripts
./setup.sh
```

Run the linter with this command:
```bash
cd api
golangci-lint run
```

### Tests
* Fill `.env.test` file
* Run these commands:
```bash
cd api
make up
go test ./... -v
```

### Ressources

Porsche-api comes with a set of 1 common resource:

`/cars` 50 cars

### Routes

**GET, POST, PUT, DELETE** HTTP methods are supported. You can use http for your requests.<br>

*GET* `/api/v1/cars`<br>

*GET* `/api/v1/cars/1`<br>

*POST* `/api/v1/cars`<br>

*PUT* `/api/v1/cars/1`<br>

*DELETE* `/api/v1/cars/1`<br>

## Client

### Set up locally
Install dependencies & run development server
```
cd client
pnpm install
pnpm dev
```

To access client development server go to http://localhost:5173
