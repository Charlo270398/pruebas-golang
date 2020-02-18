const cabeceras= {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
}

function submit(event){
    let email = document.querySelector("#email").value;
    let password = document.querySelector("#password").value;
    if(email && password){
        login(email, password);
    }
}
function login(email, password){
    const url= `/login`;
    const payload= {email: email, password: password};
    const request = {
        method: 'POST', 
        headers: cabeceras,
        body: JSON.stringify(payload),
    };
    fetch(url,request)
    .then( response => response.json() )
        .then( r => {
            if(!r.Error){
                console.log("SESION INICIADA");
            }
            else{
                alert(r.Error);
            }
        })
        .catch(err => alert(err));
}

function init () {
    document.querySelector("#submit").addEventListener('click',submit,false);
}

document.addEventListener('DOMContentLoaded',init,false);