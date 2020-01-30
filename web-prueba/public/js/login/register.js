function submit(event){
    console.log(event.target);
}
function register(nombre, apellido1, apellido2, mail, identificacion, pass){
    
    let apellidos = apellido1;
    if(apellido2){
        apellidos += " " + apellido2;
    }
    const url= `/register`;
      
    const payload= {name: nombre, surname: apellidos, mail: mail};
    const request = {
        method: 'POST', 
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