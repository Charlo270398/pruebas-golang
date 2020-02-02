const cabeceras= {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
}

function submit(event){
    let nombre = document.querySelector("#name").value;
    let apellido1 = document.querySelector("#surname1").value;
    let apellido2 = document.querySelector("#surname2").value;
    let identificacion = document.querySelector("#idnumber").value;
    let email = document.querySelector("#email").value;
    let password = document.querySelector("#pass").value;
    if(nombre && apellido1 && apellido2 && identificacion && email && password){
        register(nombre,apellido1, apellido2, email, identificacion, password);
    }
}
function register(nombre, apellido1, apellido2, email, identificacion, password){
    
    let apellidos = apellido1;
    if(apellido2){
        apellidos += " " + apellido2;
    }
    const url= `/register`;
    const payload= {nombre: nombre, identificacion:identificacion, apellidos: apellidos, email: email, password: password};
    const request = {
        method: 'POST', 
        headers: cabeceras,
        body: JSON.stringify(payload),
    };
    fetch(url,request)
    .then( response => response.json() )
    .then( r => {
        if(!r.error){

        }
        else{
            alert(r.error);
        }
    })
    .catch(err => alert(err));
}
function init () {
    document.querySelector("#submit").addEventListener('click',submit,false);
}

document.addEventListener('DOMContentLoaded',init,false);