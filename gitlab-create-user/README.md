# Gitlab-user-create

This function allow you to automatically create a new user using Gitlab API.

## Configuration

Create 2 new secrets using environment vars inside your YAML file

```yaml
gitlab-create-user:
  lang: go
  handler: ./gitlab-create-user
  image: jerome1337/gitlab-create-user:latest
  environment:
    GITLAB_API_URL: "YOUR_API_URL"
    GITLAB_API_TOKEN: "YOUR_API_TOKEN"
```

Or using `faas-cli` to allow this function to get your Gitlab API URL and token

```bash
$ faas-cli secret create GITLAB_API_URL --from-literal="YOUR_API_URL"

$ faas-cli secret create GITLAB_API_TOKEN --from-literal="YOUR_API_TOKEN"
```

## Usage

### JSON request body sample

```json
{
  "Username":"foo",
  "Name": "foo",
  "Email":"foo@gmail.com"
}
```

### Response

```json
{
  "Code": 200,
  "Message": "User foo is now created"
}
```

