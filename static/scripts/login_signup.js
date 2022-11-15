
function signup() {
    let username = document.getElementById("username")
    let password = document.getElementById("password")
    let data = {
        Username: username.value,
        Password: password.value
    }
    fetch('/signup', {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: 'POST',
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            alert(data)
            document.location.href = '/login.html'
        })
    }).catch((error) => {
        console.log(error)
    })
}
function login() {
    let username = document.getElementById("username")
    let password = document.getElementById("password")
    let data = {
        Username: username.value,
        Password: password.value
    }
    fetch('/login', {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: 'POST',
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data)
            localStorage.setItem('LoggedIn', 'true')
            localStorage.setItem('Id', result.Id)
            alert(data)
            document.location.href = '/'
        }).catch((error) => {
            console.log(error)
            alert("Please try again")
            username.value = ""
            password.value = ""

        })
    }).catch((error) => {
        console.log(error)

    })
}