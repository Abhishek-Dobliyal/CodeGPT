# CodeGPT
A utility to assist and automate coding tasks.

![](demo/codegpt_demo.gif)

> Note: Demo of U.I (and not the whole working)

## Installations and Dependencies

- [GoLang](https://go.dev/)
- [NodeJS](https://nodejs.org/en/)
- [GoFiber](https://gofiber.io/)
- [VueJS](https://vuejs.org/)
- [OpenAI API Key](https://platform.openai.com/account/api-keys)

## Usage

### Backend

- Navigate to the `backend` directory. Open Terminal/ Command Prompt and type in:

```bash
go mod tidy
```

- Plug in your `OpenAI API Key` inside the `.env` file

- Once finished, type in:

```bash
go run main.go
```

- The above command will result in a output as below:

```bash
 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.43.0                   │ 
 │               http://127.0.0.1:5000               │ 
 │       (bound on host 0.0.0.0 and port 5000)       │ 
 │                                                   │ 
 │ Handlers ............. 3  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 86476 │ 
 └───────────────────────────────────────────────────┘ 
```

### Frontend

- Open Terminal/ Command Prompt, then navigate to the `frontend` directory and type in:

```bash
npm install
```

- Once everything is setup, type in:

```bash
np run serve
```

- The above command should result in something like this:

```bash
> frontend@0.1.0 serve
> vue-cli-service serve

 INFO  Starting development server...
 DONE  Compiled successfully in 7988ms

  App running at:
  - Local:   http://localhost:8080/ 
  - Network: http://192.168.134.61:8080/
```

## Note

- Kindly do not move, delete, rename or modify any files (unless you know what you are doing).
- Feel free to make enhancements and raise a PR.

## Todo

- [ ] Wrap up backend
- [ ] Include more features
- [ ] Work on issues (if any)
- [ ] Deployment
