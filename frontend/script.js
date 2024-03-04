const container = document.getElementById('container');
const registerBtn = document.getElementById('register');
const loginBtn = document.getElementById('login');
const signupBtn = document.getElementById("submit_signup")
const signinBtn = document.getElementById('submit_signin')

async function sendData(path, name, email, password) {
    var body = JSON.stringify({ 'name':name, 'email':email, 'password':password })
    const response = await fetch(`https://mongo.lone-faerie.xyz/api/user/${path}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: body,
    });
    if (!response.ok) {
        throw Error
    }
    console.log(body)
}

signupBtn.addEventListener("click", () => {
    let name = document.getElementById("inputname").value
    let email = document.getElementById("inputemail").value
    let password = document.getElementById("inputpassword").value
    let calpolyTag = "@calpoly.edu"

    console.table({
        name,
        email,
        password,
    })

    // if (email.includes(calpolyTag)) {
    //     sendData("signup", name, email, password)
    // } else {
    //     console.log('Not a valid email')
    // }

    sendData("signup", name, email, password)
    .then(() => {
        window.location.replace("http://127.0.0.1:5500/frontend/map-demo.html");
    }).catch(() => {

    });
})

signinBtn.addEventListener("click", () => {
    let email = document.getElementById("returningEmail").value
    let password = document.getElementById("returningPassword").value

    console.table({
        email,
        password
    })

    sendData("signin", "", email, password)
    .then(() => {
        window.location.replace("http://127.0.0.1:5500/frontend/map-demo.html");
    }).catch(() => {

    });
})

registerBtn.addEventListener('click', () =>{
    container.classList.add("active")
});


loginBtn.addEventListener('click', () => {
    container.classList.remove("active");
});