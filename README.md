# Draft..!!!

# Introduction

As JWT token is one of the most famous ways of authentication. So in this post, we're quickly going to implement JWT with Go+React. We also protect our Token protected from Attacks like XSS & CSRF in simple but effective ways. This post requires basic knowledge about React so you must check-out my previous Post on [REST server with Go in 5 minutes🔗](https://dev.to/kushagra_mehta/rest-server-with-go-in-5-minutes-3n8l).

## Final Result

[Hosted on Heroku🔗](https://quiet-bastion-73016.herokuapp.com/)

# Code

Starter Code Link is 👉🏻[@JWT-with-React-Go](https://github.com/KushagraMehta/JWT-with-React-Go/tree/Starter-Code).

> In this post we're not going to focus on data validation, Error handling or somthing else as we're going to focus on securely and quickly building JWT Authentication.

# **Let's dive into the code**

![Let's dive into the code](https://media.giphy.com/media/LmNwrBhejkK9EFP504/giphy.gif)

## GoLang 🏗

Before starting anything first understand the folder structure. The folder structure can be overwhelming at first look but it helps us to organize code better.

```
├── frontend
└── backend
    ├── auth
    │   └── token.go            All the Code related Token creation/Validation/Extraction
    ├── controller
    │   └── controller.go       Contains controller handling users' request and returns a response
    ├── middleware
    │   └── middleware.go       Basic middlerware for logging and User authentication
    ├── go.mod                  the module's module path and its dependency requirements
    ├── go.sum                  Containing cryptographic hashes of the content of specific module versions
    └── main.go                 Entry file and contain all the routes
```

---

> Go dependance of project are [jwt-go](https://github.com/dgrijalva/jwt-go), [uuid](https://github.com//google/uuid), [mux](https://github.com//gorilla/mux) and [crypto](https://pkg.go.dev/golang.org/x/crypto)
>
> All the dependencies are close to standard one so it'll not be hard to understand and I'll my best to make your ride smooth by explaining everything.

---

### main.go

Let's start by defining routes for the API in the `main.go`. First, we initialize a new Router and register Handlers to it. All the handlers are defined in _[controller.go](https://github.com/KushagraMehta/JWT-with-React-Go/blob/main/backend/controller/contoller.go#L23-L94)_.

```go
router := mux.NewRouter()

router.HandleFunc("/signup", handler.PostUser).Methods(http.MethodPost)
router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
router.HandleFunc("/logout", handler.Logout).Methods(http.MethodGet)

secure := router.PathPrefix("/auth").Subrouter()
secure.HandleFunc("/", handler.Auth).Methods(http.MethodGet)
secure.HandleFunc("/user", handler.GetUser).Methods(http.MethodGet)
secure.Use(middleware.Auth) //Adding middleware for '/auth/*' route
```

1. `/signup` handler for User Signup with `POST` method.
2. `/login` handler for User Login with `POST` method.
3. `/logout` handler to Logout User with `GET` method.
4. `/auth` creating sub-route with authentication middleware
   - `/` handler so that client can check user authentication with `GET` method.
   - `/user` handler for user can access the user data with `GET` method.

### controller.go

First, create `User` structure where we store user details.

```go
type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}
type Handler struct {
	UserData map[uuid.UUID] User
}
```

```go
func (h *Handler) PostUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var tmpUser auth.User
	json.Unmarshal(reqBody, &tmpUser)

	tmpUser.Password, _ = auth.Hash(tmpUser.Password)
	userID := uuid.New()
	h.UserData[userID] = tmpUser
	generatedTokens := auth.CreateToken(userID, tmpUser)

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    generatedTokens,
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7Days expire date
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
	middleware.JSON(w, http.StatusCreated, "OK")
}
```

### token.go

### middleware.go

## React ⚛

In starter code all the user interface is preset. So first thought of the files and understand their significance.

```
├── backend
└── frontend
    ├── API.js          All the API calling will reside here
    ├── App.js          App component which acts as a container for all other components
    ├── App.scss        SCSS file for whole Site.
    ├── index.js        Render all React element into the DOM
    ├── Login.js        Component for Login & SignUp
    └── UserDetail.js   Dashboard to show user detail after successful login
```

### First Step **Sign-Up**

![Sign-up Window](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/m9s0sr27a9nwstjgvx30.png)
Whenever the user clicks on the **SignUp** button, it'll call the _signUpAPI_ function with user's data and some setAuth func(We'll discuss about this later in the post).

```jsx
<button className="submit-btn" onClick={() => signUpAPI(userData, setAuth)}>
  Sign up
</button>
```

**signUpAPI** is a Async function(Want to learn about [Async](https://javascript.info/async) and [Fetch](https://javascript.info/network) best resource👉🏻[javascript.info🔥](https://javascript.info/)). Fetch will call `/signup` with method `POST`. We converts a JavaScript value to a JSON string.

```javascript
export const signUpAPI = async (data, setAuth) => {
  const response = await fetch("/signup", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    const message = `An error has occured: ${response.status}`;
    alert(message);
    return;
  }
  setAuth(true);
};
```

[XSS attack](<https://www.vpnmentor.com/blog/top-10-common-web-attacks/#:~:text=7.%20Cross-Site%20Scripting%20(XSS),users%E2%80%99%20sessions.>)
