"use strict";

const loginRegEx = /(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)])/;

// submitForm.addEventListener('submit', async function(e) {
//     e.preventDefault();
//     try {
//         const response = await axios.post('/user/login',{
//             username: username.value,
//             password: password.value
//         }).then(
//             () => {
//                 location.replace('/transferts')
//             }
//         )
//     } catch (error) {
//         console.log(error);
//     }
// })

function displayErr(text) {
    const status = document.getElementById("error_msg");
    status.innerHTML = text;
}

function verifyCreds() {
    let login = document.getElementById("login").value;
    let password = document.getElementById("password").value;
    let first_name = document.getElementById("first_name").value;
    let last_name = document.getElementById("last_name").value;
    let grade = document.getElementById("grade").value;
    let school = document.getElementById("school").value;


    if (login === "") {
        displayErr("Email not specified");
    } else if (password === "") {
        displayErr("Password not specified");
    } else if (!loginRegEx.test(login)) {
        displayErr("Please enter a valid login");
    } else {
        let resp = obtainToken(login, password, first_name, last_name, grade, school);
        if (resp["status"] !== "ok") {
            console.log(resp);
            let payload = "";
            if (typeof resp["detail"] !== "object") {
                payload += resp["detail"].charAt(0).toUpperCase() + resp["detail"].slice(1);
            } else {
                for (let elem of resp["detail"]) {
                    if (elem["Field"] === undefined || elem["Message"] === undefined) {
                        continue;
                    }
                    payload += `${elem["Field"]}: ${elem["Message"]}<br>`;
                }
            }
            displayErr(payload);
        }
    }
}

function obtainToken(login, password, first_name, last_name, grade, school) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:1323/api/v1/register", false);
    xhr.setRequestHeader("Content-Type", "application/json;charset=utf-8");
    xhr.send(JSON.stringify({
        "Login": login,
        "Password": password,
        "FirstName": first_name,
        "LastName": last_name,
        "Grade": grade,
        "School": school
    }));

    if (xhr.status === 200) {
        location.replace("/profiles");
    }

    return JSON.parse(xhr.responseText);
}
